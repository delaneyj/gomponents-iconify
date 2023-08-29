package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.959 17.678H9.403a4.9 4.9 0 0 0-4.663 6.403l3.976 12.333a4.715 4.715 0 0 0 4.488 3.27h22.555a4.9 4.9 0 0 0 4.663-6.404l-3.975-12.333a4.715 4.715 0 0 0-4.488-3.268Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.696 39.683A6.804 6.804 0 0 0 43.5 32.88V18.258a6.804 6.804 0 0 0-6.804-6.804H24.731a4.45 4.45 0 0 1-2.478-.754l-2.432-1.63a4.45 4.45 0 0 0-2.478-.753h-2.97a5.464 5.464 0 0 0-3.775 1.514h0a4.932 4.932 0 0 0-1.508 3.963l.314 3.885"/>`),
		g.Group(children),
	)
}