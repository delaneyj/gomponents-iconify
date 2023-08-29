package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverlappingWhiteSquares(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiOverlappingWhiteSquares0" d="M60 12H28v32h32V12Z"/></defs><g fill="#fff"><path d="M44 28H12v32h32V28Z"/><use href="#openmojiOverlappingWhiteSquares0"/></g><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><use href="#openmojiOverlappingWhiteSquares0"/><path stroke-linecap="round" d="M23.5 29H11v32h32V49"/></g>`),
		g.Group(children),
	)
}