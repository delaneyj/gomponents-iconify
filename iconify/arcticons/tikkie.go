package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tikkie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.52 21.52 0 0 0 5.15 34.36L2.5 45.5l11.14-2.65A21.5 21.5 0 1 0 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.11 31.71a7.81 7.81 0 0 1-6 2.79h0a7.83 7.83 0 0 1-7.83-7.82v-5.36a7.83 7.83 0 0 1 7.83-7.82h0a7.8 7.8 0 0 1 6 2.82m-17.56 4.84h9.88m-9.88 5.73h9.88"/>`),
		g.Group(children),
	)
}