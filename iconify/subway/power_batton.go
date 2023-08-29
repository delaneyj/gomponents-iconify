package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerBatton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334.5 53.1v64c58.3 29 98.5 89 98.5 158.6c0 97.9-79.3 177.2-177.2 177.2S78.5 373.6 78.5 275.7c0-69.6 40.2-129.6 98.5-158.6v-64C85.2 85.6 19.4 172.8 19.4 275.7C19.4 406.2 125.2 512 255.7 512S492 406.2 492 275.7c0-102.9-65.8-190.1-157.5-222.6zm-78.8 143.8c21.7 0 39.4-17.7 39.4-39.4V39.4c0-21.7-17.6-39.4-39.4-39.4s-39.4 17.7-39.4 39.4v118.2c0 21.7 17.7 39.3 39.4 39.3z"/>`),
		g.Group(children),
	)
}