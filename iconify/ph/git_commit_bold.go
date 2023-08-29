package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommitBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 116h-57.21a60 60 0 0 0-117.58 0H12a12 12 0 0 0 0 24h57.21a60 60 0 0 0 117.58 0H244a12 12 0 0 0 0-24Zm-116 48a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}