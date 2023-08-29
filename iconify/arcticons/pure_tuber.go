package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PureTuber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.303 37.463l25.069-6.172a4.109 4.109 0 0 0 2.784-5.63l-3.378-7.753a7.91 7.91 0 0 0-6.405-4.705l-25.86-2.785a2.72 2.72 0 0 0-2.93 3.367l5.1 20.296a4.65 4.65 0 0 0 5.62 3.382Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.977 18.577l11.549 3.31l-9.05 7.024"/>`),
		g.Group(children),
	)
}