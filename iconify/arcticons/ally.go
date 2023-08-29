package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.144 42.83A18.764 18.764 0 0 1 5.532 22.94a18.656 18.656 0 0 1 18.78-17.767a18.687 18.687 0 0 1 18.18 18.407"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.856 32.37a7.678 7.678 0 0 1-6.97-8.074a7.611 7.611 0 0 1 7.333-7.695a7.611 7.611 0 0 1 7.334 7.694"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 23.689v18.756H31.537V24.026m-9.299 18.759l8.31-10.416h-6.7"/>`),
		g.Group(children),
	)
}