package heroku

import (
	"net/http"

	"github.com/bgentry/heroku-go"
	"github.com/remind101/empire/empire"
	"golang.org/x/net/context"
)

type Dyno heroku.Dyno

func newDyno(j *empire.ProcessState) *Dyno {
	return &Dyno{
		Command:   string(j.Command),
		Name:      string(j.Name),
		State:     j.State,
		UpdatedAt: j.UpdatedAt,
	}
}

func newDynos(js []*empire.ProcessState) []*Dyno {
	dynos := make([]*Dyno, len(js))

	for i := 0; i < len(js); i++ {
		dynos[i] = newDyno(js[i])
	}

	return dynos
}

type GetProcesses struct {
	*empire.Empire
}

func (h *GetProcesses) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	a, err := findApp(ctx, h)
	if err != nil {
		return err
	}

	// Retrieve job states
	js, err := h.JobStatesByApp(a)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	return Encode(w, newDynos(js))
}
