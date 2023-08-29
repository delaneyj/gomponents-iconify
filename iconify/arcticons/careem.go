package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Careem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M12.652 4.328a4.679 4.679 0 0 1 5.567 2.436a4.728 4.728 0 0 1-.488 4.872a4.74 4.74 0 0 1-5.707 1.322a4.674 4.674 0 0 1-2.504-3.2a4.48 4.48 0 0 1 .417-3.133a4.898 4.898 0 0 1 2.715-2.297Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" d="M28.936 4.6a19.91 19.91 0 0 1 12.474 8.819m1.484 18.319a19.886 19.886 0 0 1-37.787-3.706"/>`),
		g.Group(children),
	)
}