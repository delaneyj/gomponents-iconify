package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M454.771 10.997c-76.209-40.113-226.406 37.395-269.553 105.63c-19.222 30.534-17.862 52.538-17.862 65.022v667.874L730.602 1200l105.917-57.833V491.739L258.215 159.706c31.033-39.057 100.827-86.683 153.16-67.555l515.104 275.498l.001 724.58l106.184-57.936V309.728L454.771 10.997z"/>`),
		g.Group(children),
	)
}