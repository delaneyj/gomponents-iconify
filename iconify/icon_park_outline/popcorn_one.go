package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopcornOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M7 16h34l-7 28H14L7 16Zm13 0v28m8-28v28m5-35a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 33 9Zm-9 0a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 24 9Zm-9 0a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 15 9Z"/><path d="M22.874 9a4 4 0 1 0-7.748 0m17.748 0a4 4 0 1 0-7.748 0M16 16h16M16 44h16"/></g>`),
		g.Group(children),
	)
}