package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 58H80a70 70 0 0 0 0 140h96a70 70 0 0 0 0-140Zm0 128H80a58 58 0 0 1 0-116h96a58 58 0 0 1 0 116Zm0-96a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}