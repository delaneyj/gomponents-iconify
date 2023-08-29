package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torrdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.753 9.149l-4.886-4.543m6.559 6.773l1.417 1.714m-.533 30.301a3.727 3.727 0 1 0-7.453 0h7.454Z"/><ellipse cx="24" cy="30.276" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.888" ry="10.882"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.379 22.04c1.227 1.227-.075 5.143-2.012 8.702s-4.659 7.882-6.354 6.373s-1.472-8.087.633-11.497s6.26-5.05 7.733-3.578ZM27.215 7.653h0c-.352-.11-1.461-.485-3.215-.485c-1.564 0-2.821.363-3.215.485h0c-4.42 1.37-7.63 5.49-7.63 10.36c0 .393.024.78.065 1.162c.724.14 1.471.22 2.237.22c3.37 0 6.4-1.426 8.543-3.697a11.713 11.713 0 0 0 10.78 3.477c.041-.382.064-.77.064-1.162c0-4.87-3.21-8.99-7.629-10.36Zm3.032 1.496l4.886-4.543m-6.559 6.773l-1.417 1.714m.533 30.301a3.727 3.727 0 1 1 7.453 0h-7.454Zm3.931-21.354c-1.227 1.227.075 5.143 2.012 8.702s4.659 7.882 6.354 6.373s1.472-8.087-.633-11.497s-6.26-5.05-7.733-3.578ZM24 19.394v21.764m0-25.46V7.177"/>`),
		g.Group(children),
	)
}