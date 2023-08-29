package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minutemailer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.187 19.181L24 4.755L0 12.386l9.196 1.963l.043 4.896l2.759-2.617l-2.147-2.076l7.336 4.63z"/>`),
		g.Group(children),
	)
}