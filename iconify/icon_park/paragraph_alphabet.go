package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphAlphabet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 9H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 19H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 29H42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 39H42"/><path fill="#2F88FF" d="M11 9L7 17H15L11 9Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 19L7 17M16 19L15 17M7 17L11 9L15 17M7 17H15"/></g>`),
		g.Group(children),
	)
}