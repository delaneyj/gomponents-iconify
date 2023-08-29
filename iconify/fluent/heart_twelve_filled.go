package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.656 2.737a2.394 2.394 0 0 0-3.447-.01c-.95.975-.945 2.559.01 3.537l3.53 3.623c.146.15.384.15.53 0l3.513-3.602a2.548 2.548 0 0 0-.01-3.535a2.395 2.395 0 0 0-3.45-.009l-.336.345l-.34-.35Z"/>`),
		g.Group(children),
	)
}