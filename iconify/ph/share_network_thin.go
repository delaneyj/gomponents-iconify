package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareNetworkThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 164a36 36 0 0 0-27.92 13.3L96.25 144a35.92 35.92 0 0 0 0-32l51.83-33.3a35.93 35.93 0 1 0-4.33-6.7l-51.83 33.3a36 36 0 1 0 0 45.4l51.83 33.3A36 36 0 1 0 176 164Zm0-136a28 28 0 1 1-28 28a28 28 0 0 1 28-28ZM64 156a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm112 72a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/>`),
		g.Group(children),
	)
}