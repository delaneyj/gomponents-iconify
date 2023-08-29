package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.929 21.482l5.846-5.846a1.999 1.999 0 0 0 1.932-3.346a2 2 0 0 0-3.346 1.932l-5.846 5.846l-1.06-1.06c2.828-3.3 3.888-6.954 5.302-13.082l6.364-.707l5.657 5.657l-.707 6.363c-6.128 1.415-9.782 2.475-13.081 5.304l-1.061-1.06ZM16.596 2.037l6.347 6.346a.5.5 0 0 1-.277.848l-1.474.23l-5.656-5.657l.212-1.485a.5.5 0 0 1 .848-.282Z"/>`),
		g.Group(children),
	)
}