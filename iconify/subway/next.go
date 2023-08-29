package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 381.4l237.7-118.9L0 143.6v237.8zm237.7-118.9v118.9l237.7-118.9l-237.7-118.9v118.9zm237.7-118.9v237.8H512V143.6h-36.6z"/>`),
		g.Group(children),
	)
}