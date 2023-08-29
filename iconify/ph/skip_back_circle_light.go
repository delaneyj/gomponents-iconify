package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm34.91-135.25a6 6 0 0 0-6.09.16L102 117.17V88a6 6 0 0 0-12 0v80a6 6 0 0 0 12 0v-29.17l54.82 34.26A6 6 0 0 0 166 168V88a6 6 0 0 0-3.09-5.25ZM154 157.17L107.32 128L154 98.83Z"/>`),
		g.Group(children),
	)
}