package tops

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"log"
	"net/http"
	"sync"
	"text/tabwriter"
	"time"
)

type ProcInfo struct {
	PID           int32
	MemoryPercent float32
	MemoryInfo    *process.MemoryInfoStat
	CPUPercent    float64
	Command       string
}

type result struct {
	out []byte
	err error
}

func HandlerTops(res http.ResponseWriter, req *http.Request) {
	pids, err := process.Pids()
	if err != nil {
		handleError(res, err)
		return
	}
	procs := make([]*ProcInfo, len(pids))
	wg := sync.WaitGroup{}
	wg.Add(len(pids))
	for ix := range pids {
		proc, err := process.NewProcess(pids[ix])
		if err != nil {
			handleError(res, err)
			return
		}
		/*
			// This is apparently not implemented yet...
			if running, err := proc.IsRunning(); err != nil {
				handleError(res, err)
				return
			} else if !running {
				continue
			}
		*/
		go func(i int) {
			var err error
			p := &ProcInfo{}
			p.PID = pids[i]
			if p.Command, err = proc.Cmdline(); err != nil {
				log.Printf("Error getting Command Line: %v", err)
			}
			if p.MemoryInfo, err = proc.MemoryInfo(); err != nil {
				log.Printf("Error getting memory info: %v", err)
			}
			if p.MemoryPercent, err = proc.MemoryPercent(); err != nil {
				log.Printf("Error getting Memory: %v", err)
			}
			if p.CPUPercent, err = proc.Percent(100 * time.Millisecond); err != nil {
				log.Printf("Error getting CPU: %v", err)
			}
			if len(p.Command) > 0 {
				procs[i] = p
			}
			wg.Done()
		}(ix)
	}
	wg.Wait()

	printProcs(procs, res)
	/*
		b, err := json.Marshal(procs)
		if err != nil {
			handleError(res, err)
			return
		}
		res.Write(b)
	*/
}

func handleError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(err.Error()))
}

func printProcs(procs []*ProcInfo, res http.ResponseWriter) {
	res.WriteHeader(http.StatusOK)
	const formatHeader = "%v\t%v\t%v\t%s\n"
	const formatBody = "%d\t%.2f\t%.2f\t%s\n"
	w := tabwriter.NewWriter(res, 0, 8, 2, ' ', 0)
	fmt.Fprintf(w, formatHeader, "PID", "CPU", "Memory", "Command")
	fmt.Fprintf(w, formatHeader, "-----", "------", "-----", "----")
	for ix := range procs {
		proc := procs[ix]
		if proc == nil {
			continue
		}
		fmt.Fprintf(w, formatBody, proc.PID, proc.CPUPercent, proc.MemoryPercent, proc.Command)
	}
	w.Flush()
}
