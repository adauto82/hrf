package cmd

type ChanWriter struct {
	ch chan []byte
}

func NewChanWriter() *ChanWriter {
	return &ChanWriter{make(chan []byte, 1024)}
}

func (w *ChanWriter) Chan() <-chan []byte {
	return w.ch
}

func (w *ChanWriter) Write(p []byte) (int, error) {
	w.ch <- p
	return len(p), nil
}

func (w *ChanWriter) Close() error {
	close(w.ch)
	return nil
}
