package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm45.32-143.37l-64 32a6 6 0 0 0-2.69 2.69l-32 64A6 6 0 0 0 80 182a6.06 6.06 0 0 0 2.68-.63l64-32a6 6 0 0 0 2.69-2.69l32-64a6 6 0 0 0-8.05-8.05Zm-33.79 64.9l-46.11 23.05l23-46.11l46.11-23Z"/>`),
		g.Group(children),
	)
}