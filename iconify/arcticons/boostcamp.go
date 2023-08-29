package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boostcamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.83 39.52l6.637-3.99l-.133 7.97l-6.504-3.98Zm10.239-3.999l2.714 4.605l1.035-4.624l-3.75.02ZM15.78 24.602l-1.957-3.999l14.828-8.366L15.78 24.602Zm13.835-15.71l2.033 1.193l2.888-2.54l-.424 3.943l2.103 1.22l-.134-8.208l-6.466 4.391ZM11.784 31.82l14.848-.004l5.275-19.318l-20.123 19.323Zm17.696.044l4.692-17.281l-.074 17.274l-4.618.007Z"/>`),
		g.Group(children),
	)
}