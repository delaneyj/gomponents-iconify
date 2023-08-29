package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTurnOn0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round"><path fill="#fff" stroke-width="4" d="M21 43c-4.726-1.767-8.668-7.815-10.64-11.357c-.852-1.53-.403-3.408.964-4.502a3.83 3.83 0 0 1 5.1.283L18 29V17.5a2.5 2.5 0 0 1 5 0v6a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v2a2.5 2.5 0 0 1 5 0v7.868c0 1.07-.265 2.128-.882 3.003C36.095 39.82 34.255 42.034 32 43c-3.5 1.5-6.63 1.634-11 0Z"/><path stroke-width="3" d="m12 9l3 3m4-8l1 6m7-2l-2 3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTurnOn0)"/>`),
		g.Group(children),
	)
}