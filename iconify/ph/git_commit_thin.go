package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommitThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 124h-68.17a52 52 0 0 0-103.66 0H8a4 4 0 0 0 0 8h68.17a52 52 0 0 0 103.66 0H248a4 4 0 0 0 0-8Zm-120 48a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44Z"/>`),
		g.Group(children),
	)
}