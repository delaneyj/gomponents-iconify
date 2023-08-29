package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleBikeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 52a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm104 124a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm-16 0a24 24 0 1 0-24 24a24 24 0 0 0 24-24Zm-24-64a8 8 0 0 0-8-8h-36.69l-29.65-29.66a8 8 0 0 0-11.32 0l-32 32a8 8 0 0 0 0 11.32L120 155.31V200a8 8 0 0 0 16 0v-48a8 8 0 0 0-2.34-5.66L99.31 112L120 91.31l26.34 26.35A8 8 0 0 0 152 120h40a8 8 0 0 0 8-8ZM96 176a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm-16 0a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z"/>`),
		g.Group(children),
	)
}