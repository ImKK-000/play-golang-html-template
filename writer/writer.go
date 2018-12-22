package writer

type Writer struct {
	Data []byte
}

func (writer *Writer) Write(bytes []byte) (int, error) {
	writer.Data = append(writer.Data, bytes...)
	return 0, nil
}
