package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketsAngleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M90.12 46.44L38.22 128l51.9 81.56a12 12 0 1 1-20.24 12.88l-56-88a12 12 0 0 1 0-12.88l56-88a12 12 0 0 1 20.24 12.88Zm152 75.12l-56-88a12 12 0 1 0-20.24 12.88l51.9 81.56l-51.9 81.56a12 12 0 1 0 20.24 12.88l56-88a12 12 0 0 0 0-12.88Z"/>`),
		g.Group(children),
	)
}