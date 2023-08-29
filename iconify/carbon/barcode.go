package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 6h2v22H2zm10 0h2v20h-2zM6 6h4v20H6zm10 0h4v20h-4zm6 0h4v20h-4zm6 0h2v22h-2z"/>`),
		g.Group(children),
	)
}