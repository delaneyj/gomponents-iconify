package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#31373D" d="M27 32a4 4 0 0 1-4 4H13a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v28z"/><path fill="#77B255" d="M17.999 25a4 4 0 1 1 0 8a4 4 0 0 1 0-8z"/><path fill="#FFCC4D" d="M17.999 14a4 4 0 1 1 0 8a4 4 0 0 1 0-8z"/><path fill="#DD2E44" d="M17.999 3a4 4 0 1 1 0 8a4 4 0 0 1 0-8z"/>`),
		g.Group(children),
	)
}