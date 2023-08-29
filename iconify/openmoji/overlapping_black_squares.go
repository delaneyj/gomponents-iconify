package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverlappingBlackSquares(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiOverlappingBlackSquares0" d="M60 12H28v32h32V12Z"/></defs><path stroke="#000" stroke-width="2" d="M44 28H12v32h32V28Z"/><use href="#openmojiOverlappingBlackSquares0"/><g fill="#3F3F3F"><path stroke="#3F3F3F" stroke-width="2" d="M44 28H12v32h32V28Z"/><use href="#openmojiOverlappingBlackSquares0"/></g><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M24.5 28H12v32h32V48"/><use href="#openmojiOverlappingBlackSquares0"/></g>`),
		g.Group(children),
	)
}