package scheduler

type Options struct {
	StoragePath string
}

type SchedulerOption func(*Options)

func WithStoragePath(path string) SchedulerOption {
	return func(o *Options) {
		o.StoragePath = path
	}
}
