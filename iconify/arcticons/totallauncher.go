package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Totallauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.858 18.803V8.836a2.181 2.181 0 0 1 2.185-2.187h30.272A2.181 2.181 0 0 1 43.5 8.836V20.15a2.181 2.181 0 0 1-2.185 2.187H36.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.685 18.963h21.977a2.181 2.181 0 0 1 2.185 2.187v1.546a2.181 2.181 0 0 1-2.185 2.187H6.685A2.181 2.181 0 0 1 4.5 22.696V21.15a2.181 2.181 0 0 1 2.185-2.187Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.608 16.762a2.511 2.511 0 0 1 2.516-2.518h15.384a2.511 2.511 0 0 1 2.516 2.518v15.396a2.511 2.511 0 0 1-2.516 2.518H18.124a2.511 2.511 0 0 1-2.516-2.518V24.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.354 34.78v4.384a2.181 2.181 0 0 1-2.185 2.187h-10.19a2.181 2.181 0 0 1-2.186-2.187V25.241"/><path fill="none" stroke="currentColor" stroke-miterlimit="8.8" d="M15.608 16.762v1.92"/>`),
		g.Group(children),
	)
}