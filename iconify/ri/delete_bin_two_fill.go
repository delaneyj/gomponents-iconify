package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteBinTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3h5v2h-2v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H2V6h5Zm6.414 8l1.768-1.768l-1.414-1.414L12 12.586l-1.768-1.768l-1.414 1.414L10.586 14l-1.768 1.768l1.414 1.414L12 15.414l1.768 1.768l1.414-1.414L13.414 14ZM9 4v2h6V4H9Z"/>`),
		g.Group(children),
	)
}