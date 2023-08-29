package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Evie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.51a21.5 21.5 0 0 0-12.61 38.92c.44-10 .83-20 3-29.12c1.23.75 2.35 2.25 3.79 2.63a7.44 7.44 0 0 1 .53-2.53c1.77 1.1 4.26 3.56 6 3.25a22.57 22.57 0 0 1 10.37.44c1.38 8.72-2.62 12.65-7.49 15.82c4.68.67 8.43 3.5 11.89 7A21.5 21.5 0 0 0 24 2.51Zm.34 17.43a2.24 2.24 0 1 0 2.27 2.23a2.23 2.23 0 0 0-2.27-2.23Zm0 0"/>`),
		g.Group(children),
	)
}