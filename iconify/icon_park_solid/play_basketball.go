package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayBasketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPlayBasketball0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-miterlimit="2" stroke-width="4" d="M32 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="m23 40l8.11-2.06c.78-.19 1.02-1.19.42-1.72L23 29l4-8l-10.41-3.74c-.5-.18-.9-.54-1.13-1.02L11 8"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="m23 29l-6.97 8.79c-.21.25-.47.45-.77.57L5 42m22-21l9.9 2.79c.47.14.88.44 1.14.85L42 31"/><path fill="#fff" d="M18 8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPlayBasketball0)"/>`),
		g.Group(children),
	)
}