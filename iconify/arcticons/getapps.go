package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Getapps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.532 17.59L24.484 36.702m6.639 6.769L11.784 23.907m12.152-12.216l-19.468 19.5m12.828-26.72l19.661 19.694m-12.973-5.577l.032 10.83l5.173-5.092"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.048 18.588l-.032 10.83l-5.173-5.092"/>`),
		g.Group(children),
	)
}