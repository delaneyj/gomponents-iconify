package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expenses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 34h12m-5 5l5-5l-5-5"/><path d="M43 26V10a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v28a3 3 0 0 0 3 3h20.47"/><path d="m15 15l5 6l5-6M14 27h12m-12-6h12m-6 0v12"/></g>`),
		g.Group(children),
	)
}