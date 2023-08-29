package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ancestry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.857 7.98v16.037L43.5 40.02m-.643-16.003c-6.725 0-18.044.32-18.044.32m0 0c-6.697-7.732-15.13-5.728-20.313-.51c3.586 5.666 16.558 7.531 20.313.51Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.99 24.043c.616 6.178-4.505 17.058-20.24 15.89c.903-1.233-.13-10.693 12.843-15.762"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.99 24.043c.616-6.178-4.505-17.058-20.24-15.89c1.496 1.56-1.53 11.926 12.563 16.024"/>`),
		g.Group(children),
	)
}