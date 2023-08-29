package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm0-124a40 40 0 0 0-40 40a12 12 0 0 1-24 0a64.07 64.07 0 0 1 64-64a12 12 0 0 1 0 24Zm64 40a64.07 64.07 0 0 1-64 64a12 12 0 0 1 0-24a40 40 0 0 0 40-40a12 12 0 0 1 24 0Zm-64 20a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}