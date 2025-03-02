package models

//  //
//type URLFileHandler struct {
//	file    *os.File
//	writer  *bufio.Writer
//	scanner *bufio.Scanner
//}
//
//func (p *URLFileHandler) WriteURLToFile(d *easyjson.URLFileRecord) error {
//
//	defer p.close()
//
//	err := p.writeURLToFile(d)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (p *URLFileHandler) ReadURLFromFile() (*easyjson.URLFileRecord, error) {
//
//	defer p.close()
//
//	record, err := p.readURLFromFile()
//	if err != nil {
//		return nil, err
//	}
//	return record, nil
//}
//func NewFileHandler(path string, mode int) (*URLFileHandler, error) {
//	file, err := os.OpenFile(path, mode, 0666)
//	if err != nil {
//		return nil, err
//	}
//	handler := &URLFileHandler{
//		file: file,
//	}
//	if mode == os.O_WRONLY|os.O_CREATE|os.O_APPEND {
//		handler.writer = bufio.NewWriter(file)
//	} else if mode == os.O_RDONLY|os.O_CREATE {
//		handler.scanner = bufio.NewScanner(file)
//	}
//	return handler, nil
//}
//
//func (p *URLFileHandler) writeURLToFile(d *easyjson.URLFileRecord) error {
//	if p.writer == nil {
//		_ = fmt.Errorf("FIleHandler is not initialized for writing")
//	}
//	data, err := d.MarshalJSON()
//	if err != nil {
//		return err
//	}
//
//	if _, err := p.writer.Write(data); err != nil {
//		return err
//	}
//
//	if err := p.writer.WriteByte('\n'); err != nil {
//		return err
//	}
//	return p.writer.Flush()
//}
//func (p *URLFileHandler) readURLFromFile() (*easyjson.URLFileRecord, error) {
//	if p.scanner == nil {
//		return nil, fmt.Errorf("FileHandler is not initialized for reading")
//	}
//	if !p.scanner.Scan() {
//		return nil, p.scanner.Err()
//	}
//	data := p.scanner.Bytes()
//
//	record := &easyjson.URLFileRecord{}
//	if err := record.UnmarshalJSON(data); err != nil {
//		return nil, err
//	}
//	return record, nil
//}
//func (p *URLFileHandler) close() error {
//	return p.file.Close()
//}
