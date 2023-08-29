package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverlappingWhiteAndBlackSquares(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiOverlappingWhiteAndBlackSquares0" d="M60 12H28v32h32V12Z"/></defs><use href="#openmojiOverlappingWhiteAndBlackSquares0"/><path fill="#fff" stroke="#fff" stroke-width="2" d="M44 28H12v32h32V28Z"/><path fill="#3F3F3F" d="M60 12H28v32h32V12Z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M24.5 28H12v32h32V48"/><use href="#openmojiOverlappingWhiteAndBlackSquares0"/></g>`),
		g.Group(children),
	)
}