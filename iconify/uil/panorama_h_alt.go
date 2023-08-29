package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.46 5.83A1 1 0 0 0 20.7 5h-.11A37.49 37.49 0 0 0 3.41 5H3.3a1 1 0 0 0-.76.8a35.52 35.52 0 0 0 0 12.34a1 1 0 0 0 .76.8h.11A37.62 37.62 0 0 0 12 20a37.62 37.62 0 0 0 8.59-1h.11a1 1 0 0 0 .76-.8a35.52 35.52 0 0 0 0-12.37ZM19.6 17.17a35.42 35.42 0 0 1-15.2 0a33.2 33.2 0 0 1 0-10.34a35.42 35.42 0 0 1 15.2 0a33.2 33.2 0 0 1 0 10.34Z"/>`),
		g.Group(children),
	)
}