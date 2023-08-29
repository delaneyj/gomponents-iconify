package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 76a40 40 0 1 0-40 40a40 40 0 0 0 40-40Zm-40 32a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm60 24a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 72a32 32 0 1 1 32-32a32 32 0 0 1-32 32ZM68 132a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 72a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}