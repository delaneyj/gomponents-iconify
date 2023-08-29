package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crystal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M93.365 0L0 93.741l34.258 128.376L163.011 256L256 161.882L222.117 35.388L93.365 0ZM6.776 95.247l124.988-32.754l-34.258 124.612l-90.73-91.858Z"/>`),
		g.Group(children),
	)
}