package libxml
/*
#cgo LDFLAGS: -lxml2
#cgo CFLAGS: -I/usr/include/libxml2
#include <libxml/HTMLparser.h> 
#include <libxml/HTMLtree.h>
*/
import "C"

import (
	//"unsafe"
)

type HtmlDoc struct {
	*XmlDoc
}

func ParseHtmlString(content string, url string, encoding string, opts int) *XmlDoc {
	cString := C.CString(content)
	cXmlChar := C.xmlCharStrdup(cString)
	xmlDocPtr := C.htmlReadDoc(cXmlChar, C.CString(url), C.CString(encoding), C.int(opts))
	return buildXmlDoc(xmlDocPtr)
}

func HtmlParse(content string) *XmlDoc {
	return ParseHtmlString(content, "", "",
		HTML_PARSE_RECOVER|
			HTML_PARSE_NONET|
			HTML_PARSE_NOERROR|
			HTML_PARSE_NOWARNING)
}

func ParseHtmlFile(url string, encoding string, opts int) *XmlDoc {
	return buildXmlDoc(C.htmlReadFile(C.CString(url), C.CString(encoding), C.int(opts)))
}
