package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignNoParkingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M13.292 14A8 8 0 0 1 2 2.707l3.5 3.5V12h1.283V9.164h1.674L13.292 14Zm.708-.708l-4.37-4.37C10.5 8.524 11 7.662 11 6.587c0-1.482-.955-2.584-2.538-2.584H5.5v.79L2.708 2.002A8 8 0 0 1 14 13.293Z"/><path d="M6.777 7.485v.59h.59l-.59-.59Zm1.949.535L6.777 6.07v-.966H8.27c.893 0 1.419.539 1.419 1.482c0 .769-.35 1.273-.963 1.433Z"/></g>`),
		g.Group(children),
	)
}