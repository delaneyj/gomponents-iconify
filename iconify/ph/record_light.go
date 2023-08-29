package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm0-160a70 70 0 1 0 70 70a70.08 70.08 0 0 0-70-70Zm0 128a58 58 0 1 1 58-58a58.07 58.07 0 0 1-58 58Z"/>`),
		g.Group(children),
	)
}