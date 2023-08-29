package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffdd7d" d="M59.7 13.8c1.7-5.2 1.2-14.9-16.6-3.2C25 22.5 6.2 50.3 7.6 55.4c1.1 4 17.3-5 26-17.2c.7-1 8.7 8.8 7.6 9.8C33.4 55.8 8.3 65.7 3 60.6c-6.1-5.9 16.7-39.8 40.1-53.3c6.4-3.7 25.5-12.5 16.6 6.5"/><path fill="#ffd05a" d="M60.6 49.5L46.7 48l-9.1 10.6l-2.9-13.7l-12.9-5.4l12.2-7l1.2-13.9L45.5 28l13.6-3.2l-5.7 12.7l7.2 12"/>`),
		g.Group(children),
	)
}