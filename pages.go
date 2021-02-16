package docconv

import (
	"io"
)

// ConvertPages converts a Pages file to text.
func ConvertPages(r io.Reader) (textBody string, meta map[string]string, err error) {
	panic("ConvertPages is not implemented!")

	//prefix := util.NewErrorPrefixer("ConvertPages")
	//b := &bytes.Buffer{}
	//size, err := b.ReadFrom(r)
	//if err != nil {
	//	err = prefix(err, "failed to read")
	//	return
	//}
	//
	//zr, err := zip.NewReader(bytes.NewReader(b.Bytes()), size)
	//if err != nil {
	//	err = prefix(err, "failed to unzip data")
	//	return
	//}
	//
	//for _, f := range zr.File {
	//	var r io.ReadCloser
	//	r, err = f.Open()
	//	if err != nil {
	//		err = prefix(err, "failed to open file")
	//		return
	//	}
	//	defer util.HandleCloseError(prefix(r.Close(), "failed to close file"), &err)
	//	switch {
	//	case strings.HasSuffix(f.Name, "Preview.pdf"):
	//		return ConvertPDF(r)
	//	case f.Name == "index.xml":
	//		return ConvertXML(r)
	//	case f.Name == "Index/Document.iwa":
	//		var (
	//			archiveLength int64
	//			archiveInfoData []byte
	//		)
	//		bReader := bufio.NewReader(snappy.NewReader(io.MultiReader(strings.NewReader("\xff\x06\x00\x00sNaPpY"), r)))
	//		archiveLength, err = binary.ReadVarint(bReader)
	//		if err != nil {
	//			err = prefix(err, "failed to read Document.iwa archive length")
	//			return
	//		}
	//		archiveInfoData, err = ioutil.ReadAll(io.LimitReader(bReader, archiveLength))
	//		if err != nil {
	//			err = prefix(err, "failed to read Document.iwa")
	//			return
	//		}
	//		archiveInfo := TSP.ArchiveInfo{}
	//		if err = proto.Unmarshal(archiveInfoData, &archiveInfo); err != nil {
	//			err = prefix(err, "failed to unmarshal archive data")
	//			return
	//		}
	//	}
	//}
	//return
}
