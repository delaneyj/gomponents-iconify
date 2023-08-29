package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 52H80a76 76 0 0 0 0 152h96a76 76 0 0 0 0-152Zm0 128H80a52 52 0 0 1 0-104h96a52 52 0 0 1 0 104ZM80 92a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 48a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}