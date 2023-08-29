package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsSixVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 16H64a16 16 0 0 0-16 16v192a16 16 0 0 0 16 16h128a16 16 0 0 0 16-16V32a16 16 0 0 0-16-16Zm-92 184a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-60a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-60a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm56 120a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-60a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-60a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}