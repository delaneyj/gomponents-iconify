package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Runnerup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.865 5.5a5.093 5.093 0 1 1-3.677 1.49a5.095 5.095 0 0 1 3.677-1.49Zm13.467 1.09l8.028 8.028l-3.351 3.352l-4.678-4.677l-3.744 3.747L42.89 36.341l-3.351 3.352L28.926 29.08l-3.744 3.747l6.322 6.322l-3.352 3.351l-9.673-9.673l7.096-7.098l-5.34-5.339l-7.097 7.096l-8.028-8.029l3.351-3.35l4.677 4.676L27.331 6.59Z"/>`),
		g.Group(children),
	)
}