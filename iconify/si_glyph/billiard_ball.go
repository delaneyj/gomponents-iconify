package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BilliardBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.01.057a7.94 7.94 0 1 0 .002 15.88A7.94 7.94 0 0 0 9.01.057zm-1 12.057c-2.804 0-5.076-2.283-5.076-5.099c0-2.814 2.272-5.098 5.076-5.098c2.805 0 5.078 2.283 5.078 5.098c0 2.816-2.274 5.099-5.078 5.099z"/><path d="M8 4h.938v5.938H8z"/><path d="M7 4h1.969v.938H7z"/></g>`),
		g.Group(children),
	)
}