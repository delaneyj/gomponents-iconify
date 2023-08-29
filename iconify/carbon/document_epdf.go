package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentEpdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 26v-4a2.002 2.002 0 0 0-2-2h-4a2.002 2.002 0 0 0-2 2v6a2.008 2.008 0 0 0 2 2h5v-2h-5v-2Zm-6-4h4v2h-4Z"/><path fill="currentColor" d="m25.707 9.293l-7-7A1 1 0 0 0 18 2H8a2.002 2.002 0 0 0-2 2v24a2.002 2.002 0 0 0 2 2h8v-2H8V4h8v6a2.002 2.002 0 0 0 2 2h6v4h2v-6a1 1 0 0 0-.293-.707ZM18 4.414L23.586 10H18Z"/>`),
		g.Group(children),
	)
}