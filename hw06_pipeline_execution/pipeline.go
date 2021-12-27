package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func StopOrNot(in, done In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			close(out)
			for range in {
			}
		}()
		for {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := StopOrNot(in, done)
	for _, v := range stages {
		out = v(StopOrNot(out, done))
	}
	return out
}
