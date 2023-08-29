package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuKebabHorizontalSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path stroke-width="3.75" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`),
		g.Group(children),
	)
}