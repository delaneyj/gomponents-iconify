package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="30" x="24.762" y="3.243" rx="2" transform="rotate(45 24.762 3.243)"/><path d="m38.197 16.677l4.242-4.242l-7.07-7.071l-4.244 4.242M18 21h12m-12 6h4m6 0h2"/></g>`),
		g.Group(children),
	)
}