package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ucgo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="11.339" d="M10.622 9.501v8.484a4.203 4.203 0 0 0 8.405 0V9.501m12.233 8.431v.052a4.203 4.203 0 0 1-8.404 0v-4.281A4.203 4.203 0 0 1 27.058 9.5h0a4.203 4.203 0 0 1 4.203 4.203v.052m-6.117 16.26a4.203 4.203 0 0 0-8.405 0v4.282a4.203 4.203 0 0 0 8.405 0h-4.202M33.175 38.5a4.203 4.203 0 0 1-4.202-4.202v-4.282a4.203 4.203 0 0 1 8.405 0v4.282a4.203 4.203 0 0 1-4.203 4.202Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}