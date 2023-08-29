package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.383 2.076a1 1 0 0 0-1.09.217l-26 26A1 1 0 0 0 3 30h8v-2H5.414l4-4H22a2.002 2.002 0 0 0 2-2V9.415l4-4V11h2V3a1 1 0 0 0-.617-.924zM22 22H11.414L22 11.415zm-12-4.245V10h7.755l2-2H10a2.002 2.002 0 0 0-2 2v9.755zM11 2H2v9h2V4h7V2zm10 28h9v-9h-2v7h-7v2zM4 23.755V21H2v4.754l2-1.999zM25.755 2H21v2h2.755l2-2z"/>`),
		g.Group(children),
	)
}