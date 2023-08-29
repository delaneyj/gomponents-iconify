package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#347FFB" fill-rule="nonzero"/><path fill="#FFF" d="m19.7 13.055l-2.869 7.75l-.018.047l-.526-1.055l-1.1.5l2.885-7.797l1.1-.5l.527 1.055zm7.18.183L24.012 21l-.527-1.058l-1.1.5l.067-.182l2.867-7.76l1.1-.5l.525 1.055l-.064.183zm-6.14 6.712l1.689-4.563l-1.103.5l-.524-1.057l-1.694 4.562l.525 1.058l1.107-.5zm-9.137-4.5H6.558l.86.8l-.86.813h5.04l.856-.813l-.851-.8zM5.86 18.833h8.155l-.857.807l.857.805H5.86L5 19.64l.86-.808zm2.501-6.768h8.15l-.854.808l.855.805h-8.15l-.86-.806l.86-.807z"/></g>`),
		g.Group(children),
	)
}