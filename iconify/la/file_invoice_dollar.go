package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileInvoiceDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h16v-2H8V5h10v6h6v2h2V9.6l-.3-.3l-6-6l-.3-.3H6zm14 3.4L22.6 9H20V6.4zM10 13v2h12v-2H10zm17 2v2c-1.7.3-3 1.7-3 3.5c0 2 1.5 3.5 3.5 3.5h1c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5H25v2h2v2h2v-2c1.7-.3 3-1.7 3-3.5c0-2-1.5-3.5-3.5-3.5h-1c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5H31v-2h-2v-2h-2zm-17 3v2h7v-2h-7zm9 0v2h3v-2h-3zm-9 4v2h7v-2h-7zm9 0v2h3v-2h-3z"/>`),
		g.Group(children),
	)
}