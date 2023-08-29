package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LtePlusMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 16V8h2v6h3v2H1Zm6 0v-6H5V8h6v2H9v6H7Zm5 0V8h5v2h-3v1h3v2h-3v1h3v2h-5Zm8-1v-2h-2v-2h2V9h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}