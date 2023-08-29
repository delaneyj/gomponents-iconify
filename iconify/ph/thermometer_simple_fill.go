package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 146.08V40a32 32 0 0 0-64 0v106.08a56 56 0 1 0 64 0ZM128 24a16 16 0 0 1 16 16v40h-32V40a16 16 0 0 1 16-16Z"/>`),
		g.Group(children),
	)
}