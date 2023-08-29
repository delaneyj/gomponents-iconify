package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#31373D" d="M36 23a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V13a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v10z"/><circle cx="7" cy="18" r="4" fill="#77B255"/><circle cx="18" cy="18" r="4" fill="#FFCC4D"/><circle cx="29" cy="18" r="4" fill="#DD2E44"/>`),
		g.Group(children),
	)
}