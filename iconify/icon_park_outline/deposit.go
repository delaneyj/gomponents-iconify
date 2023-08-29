package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deposit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 16v28h24V16"/><path d="m19 20l5 6l5-6M18 32h12m-12-6h12m-6 0v12M14 10h20m2 8h6V4H6v14h6"/></g>`),
		g.Group(children),
	)
}