package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 172a28 28 0 0 0-27.71 24H72a36 36 0 0 1 0-72h96a36 36 0 0 0 0-72H72a4 4 0 0 0 0 8h96a28 28 0 0 1 0 56H72a44 44 0 0 0 0 88h100.29A28 28 0 1 0 200 172Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}