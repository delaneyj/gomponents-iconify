package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Writeilypro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.31 15.93L12.24 41L5.5 42.5L7 35.76l.51-.51l28-28l.51-.51l1.76 1.76l.55-.5l2.46-2.5h0l1.71 1.71L40 9.68l-.52.51l1.8 1.81l-3 3l1.41 1.42l1.45 1.44l-4.88 4.88m1.55-14.26l1.71 1.71m-1.21 4.74l-1 1M7 35.76L12.24 41"/>`),
		g.Group(children),
	)
}