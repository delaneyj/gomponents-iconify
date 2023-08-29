package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenSourceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c5.523 0 10 4.477 10 10c0 4.13-2.504 7.676-6.077 9.201l-2.518-6.55A3 3 0 0 0 12 9a3 3 0 0 0-1.404 5.652l-2.518 6.55A10.003 10.003 0 0 1 2 12C2 6.477 6.477 2 12 2Z"/>`),
		g.Group(children),
	)
}