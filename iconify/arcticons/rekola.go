package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rekola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.504 16.25l-3.201 11.947H36.26m-10.958-.001L14.595 17.53l-2.798 10.444l4.056-15.134s-3.482-.678-4.16-.096c-.947.812-.765 2.646 1.483 2.404"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.362 26.62c1.618-6.031 7.712-6.328 10.857-4.488a7.421 7.421 0 0 1 2.765 8.89a7.296 7.296 0 0 1-13.706-.7"/><path fill="none" stroke="currentColor" d="M15.602 21.745a7.297 7.297 0 1 1-3.805-1.07m4.472-1.64l11.521-.072"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.79 15.716l4.162-.058h0"/>`),
		g.Group(children),
	)
}