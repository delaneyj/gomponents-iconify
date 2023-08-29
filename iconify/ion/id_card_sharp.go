package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M408 16H104a24 24 0 0 0-24 24v432a24 24 0 0 0 24 24h304a24 24 0 0 0 24-24V40a24 24 0 0 0-24-24Zm-61.1 296.77a43 43 0 1 1-40.71-40.71a43 43 0 0 1 40.71 40.71ZM192 64h128v32H192Zm192 384H224v-24.6c0-32.72 53.27-49.21 80-49.21s80 16.49 80 49.21Z"/>`),
		g.Group(children),
	)
}