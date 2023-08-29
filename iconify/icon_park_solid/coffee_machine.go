package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCoffeeMachine0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 42h29a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H17"/><path fill="#fff" d="M22 36c5.523 0 10-4.477 10-10H12c0 5.523 4.477 10 10 10Z"/><path stroke-linecap="round" d="M16 18v2m6-2v2m6-2v2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCoffeeMachine0)"/>`),
		g.Group(children),
	)
}