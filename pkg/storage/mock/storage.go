package mock

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/uswitch/brigade/pkg/brigade"
)

var (
	// Now is the date used in all stub date fields.
	Now = time.Now()
	// StubProject is a Project stub.
	StubProject = &brigade.Project{
		ID:           "project-id",
		Name:         "project-name",
		SharedSecret: "shared-secre3t",
		Secrets:      map[string]string{"key": "value"},
	}
	// StubWorker is a stub Worker. It is used in StubBuild, too.
	StubWorker = &brigade.Worker{
		ID:        "worker-id",
		BuildID:   "build-id",
		ProjectID: "project-id",
		StartTime: Now,
		EndTime:   Now,
		ExitCode:  0,
		Status:    brigade.JobSucceeded,
	}
	// StubBuild is a stub Build.
	StubBuild = &brigade.Build{
		ID:        "build-id",
		ProjectID: "project-id",
		Revision: &brigade.Revision{
			Commit: "commit",
		},
		Type:     "type",
		Provider: "provider",
		Payload:  []byte("payload"),
		Script:   []byte("script"),
		Worker:   StubWorker,
	}
	// StubJob is a stub Job.
	StubJob = &brigade.Job{
		ID:           "job-id",
		Name:         "job-name",
		Image:        "image",
		CreationTime: Now,
		StartTime:    Now,
		EndTime:      Now,
		ExitCode:     0,
		Status:       brigade.JobSucceeded,
	}
	// StubLogData is string data representing a log
	StubLogData = "Hello World"
)

// New returns a new Store with the default stubs.
func New() *Store {
	return &Store{
		Project: StubProject,
		Worker:  StubWorker,
		Build:   StubBuild,
		Job:     StubJob,
		LogData: StubLogData,
	}
}

// Store implements the storage.Storage interface, but returns mock data.
type Store struct {
	// Project is the project you want returned.
	Project *brigade.Project
	// Build is the build you want returned.
	Build *brigade.Build
	// Job is the job you want returned.
	Job *brigade.Job
	// Worker is the worker you want returned.
	Worker *brigade.Worker
	// LogData is the log data you want returned.
	LogData string
}

// GetProjects gets the mock project wrapped as a slice of projects.
func (s *Store) GetProjects() ([]*brigade.Project, error) {
	return []*brigade.Project{s.Project}, nil
}

// GetProject returns the Project
func (s *Store) GetProject(id string) (*brigade.Project, error) {
	return s.Project, nil
}

// GetProjectBuilds returns the mock Build wrapped in a slice.
func (s *Store) GetProjectBuilds(p *brigade.Project) ([]*brigade.Build, error) {
	return s.GetBuilds()
}

// GetBuilds returns the mock build wrapped in a slice.
func (s *Store) GetBuilds() ([]*brigade.Build, error) {
	return []*brigade.Build{s.Build}, nil
}

// GetBuild gets the mock Build
func (s *Store) GetBuild(id string) (*brigade.Build, error) {
	return s.Build, nil
}

// GetBuildJobs gets the mock job wrapped in a slice.
func (s *Store) GetBuildJobs(b *brigade.Build) ([]*brigade.Job, error) {
	return []*brigade.Job{s.Job}, nil
}

// GetWorker gets the mock worker.
func (s *Store) GetWorker(bid string) (*brigade.Worker, error) {
	return s.Worker, nil
}

// GetJob gets the mock job.
func (s *Store) GetJob(id string) (*brigade.Job, error) {
	return s.Job, nil
}

// GetJobLog gets the mock log data
func (s *Store) GetJobLog(j *brigade.Job) (string, error) {
	return s.LogData, nil
}

// GetJobLogStream gets the mock log data as a readcloser
func (s *Store) GetJobLogStream(j *brigade.Job) (io.ReadCloser, error) {
	return rc(s.LogData), nil
}

// GetWorkerLog gets the mock log data.
func (s *Store) GetWorkerLog(w *brigade.Worker) (string, error) {
	return s.LogData, nil
}

// GetWorkerLogStream gets a readcloser of the mock log data.
func (s *Store) GetWorkerLogStream(w *brigade.Worker) (io.ReadCloser, error) {
	return rc(s.LogData), nil
}

// CreateBuild fakes a new build.
func (s *Store) CreateBuild(b *brigade.Build) error {
	s.Build = b
	return nil
}

// rc wraps a string in a ReadCloser.
func rc(s string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBufferString(s))
}
