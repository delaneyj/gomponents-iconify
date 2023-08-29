package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M4 32.083c0-1.202.266-2.395.971-3.368C7.045 25.85 12.671 20 24 20c11.33 0 16.955 5.851 19.029 8.715c.705.973.971 2.166.971 3.368A7.917 7.917 0 0 1 36.083 40H11.917A7.917 7.917 0 0 1 4 32.083Z"/><path d="M12 9v4m2 9v4M36 9v4m-2 9v4M24 7v6m0 7v8m16-2.557C36.906 22.78 31.808 20 24 20s-12.906 2.779-16 5.443"/></g>`),
		g.Group(children),
	)
}