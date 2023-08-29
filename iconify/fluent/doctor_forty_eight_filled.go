package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 6A2.25 2.25 0 0 0 16 8.25V16H8.25A2.25 2.25 0 0 0 6 18.25v11.5A2.25 2.25 0 0 0 8.25 32H16v7.75A2.25 2.25 0 0 0 18.25 42h11.5A2.25 2.25 0 0 0 32 39.75V32h7.75A2.25 2.25 0 0 0 42 29.75v-11.5A2.25 2.25 0 0 0 39.75 16H32V8.25A2.25 2.25 0 0 0 29.75 6h-11.5Z"/>`),
		g.Group(children),
	)
}