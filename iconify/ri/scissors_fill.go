package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.683 7.559L12 9.875l6.374-6.374a2 2 0 0 1 2.829 0l.707.707L9.683 16.434a4 4 0 1 1-2.121-2.121l2.317-2.316L7.562 9.68a4 4 0 1 1 2.121-2.121ZM6 7.997a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm9.535-6.587l6.375 6.375l-.707.707a2 2 0 0 1-2.829 0l-4.96-4.96l2.12-2.122Z"/>`),
		g.Group(children),
	)
}