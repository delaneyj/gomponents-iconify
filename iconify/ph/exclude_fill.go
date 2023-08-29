package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcludeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 160a80 80 0 0 1-158.64 14.64a80 80 0 0 0 93.28-93.28A80 80 0 0 1 240 160Zm-80-80a80.29 80.29 0 0 1 14.64 1.36a80 80 0 1 0-93.28 93.28A80 80 0 0 1 160 80Z"/>`),
		g.Group(children),
	)
}