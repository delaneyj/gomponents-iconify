package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M462 3L290 179l-62-60l-26 25q-11-30-37.5-49T105 76q-43 0-73 29.5T2 176t29.5 70t71.5 30l125 121l143-139l-14-14l105-82V3zm-15 150L216 335l-87-77q-16 4-25 4q-37 0-62.5-25.5T16 176t25.5-60T104 91q36 0 62 25t26 60q0 18-8 36l35 57L447 35v118z"/>`),
		g.Group(children),
	)
}