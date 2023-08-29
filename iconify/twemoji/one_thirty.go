package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#99AAB5"/><circle cx="18" cy="18" r="14" fill="#E1E8ED"/><path fill="#67757F" d="M18 31a1 1 0 0 1-1-1V18a1 1 0 0 1 2 0v12a1 1 0 0 1-1 1z"/><path fill="#67757F" d="M17.999 19a1 1 0 0 1-.893-1.447l4-8a1.001 1.001 0 0 1 1.789.895l-4 8c-.176.35-.529.552-.896.552z"/>`),
		g.Group(children),
	)
}