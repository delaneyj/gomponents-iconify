package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LteMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16V8h2v6h3v2H4Zm7 0v-6H9V8h6v2h-2v6h-2Zm5 0V8h5v2h-3v1h3v2h-3v1h3v2h-5Z"/>`),
		g.Group(children),
	)
}