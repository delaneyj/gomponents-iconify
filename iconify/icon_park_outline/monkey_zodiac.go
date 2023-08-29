package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonkeyZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.593 18.134c1.34-1.464 2.915-4.977-1.508-7.32c-1.006-.65-1.585-2.314-2.514-5.368C13.718 3.33 5 3 4 14v30"/><path d="M17 25c4.582 1.673 11.033 5.254 15 16c1.533 3.754 8.507 4.87 11.5-1c.998-1.957.5-5.496-3.918-7.55C36.462 31 34 26 38.5 24c1.848-.603 3.93.08 5.5 3"/><path d="M29 35c-4.345-1.106-13.228-.481-14 9"/></g>`),
		g.Group(children),
	)
}