package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1.056 8.306l4.308-3.351l6.703 6.705l-3.35 4.308s.044-1.054-1.437-1.437c-1.39-.359-1.632-2.288-2.843-2.688c-1.211-.4-1.426-1.242-1.986-2.145S1.056 8.306 1.056 8.306zm12.925-.476l-1.146-1.146c1.841-1.603 3.279-2.902 3.537-3.162A2.032 2.032 0 0 0 13.499.649c-.258.26-1.559 1.697-3.16 3.54L9.193 3.043a1.355 1.355 0 0 0-1.916-.001l-.957.957l6.704 6.705l.958-.959a1.355 1.355 0 0 0-.001-1.915zm.265-5.053a.977.977 0 1 1 1.383-1.383a.976.976 0 0 1 0 1.383a.976.976 0 0 1-1.383 0z"/>`),
		g.Group(children),
	)
}