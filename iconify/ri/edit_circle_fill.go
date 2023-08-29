package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.626 3.128L9.29 10.463l.01 4.247l4.238-.008l7.331-7.331A9.956 9.956 0 0 1 22 11.997c0 5.522-4.477 10-10 10s-10-4.478-10-10c0-5.523 4.477-10 10-10a9.96 9.96 0 0 1 4.626 1.131Zm3.86-1.03l1.413 1.413l-9.192 9.193l-1.412.002l-.002-1.416l9.192-9.193Z"/>`),
		g.Group(children),
	)
}