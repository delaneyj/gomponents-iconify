package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 112a36 36 0 1 1 36-36a36 36 0 0 1-36 36Zm60 24a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm-120 0a36 36 0 1 0 36 36a36 36 0 0 0-36-36Z" opacity=".2"/><path d="M172 76a44 44 0 1 0-44 44a44.05 44.05 0 0 0 44-44Zm-44 28a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm60 24a44 44 0 1 0 44 44a44.05 44.05 0 0 0-44-44Zm0 72a28 28 0 1 1 28-28a28 28 0 0 1-28 28ZM68 128a44 44 0 1 0 44 44a44.05 44.05 0 0 0-44-44Zm0 72a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/></g>`),
		g.Group(children),
	)
}