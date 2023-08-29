package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m4.5 10.5l-.447-.224a.5.5 0 0 0 .67.671L4.5 10.5Zm2-4l-.224-.447a.5.5 0 0 0-.223.223L6.5 6.5Zm4-2l.447.224a.5.5 0 0 0-.67-.671l.223.447Zm-2 4l.224.447a.5.5 0 0 0 .223-.223L8.5 8.5Zm-1 5.5A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM4.947 10.724l2-4l-.894-.448l-2 4l.894.448Zm1.777-3.777l4-2l-.448-.894l-4 2l.448.894Zm3.329-2.67l-2 4l.894.447l2-4l-.894-.448ZM8.276 8.052l-4 2l.448.894l4-2l-.448-.894Z"/>`),
		g.Group(children),
	)
}