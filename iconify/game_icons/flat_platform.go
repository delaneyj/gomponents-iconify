package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlatPlatform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.5 87.9L39.7 213.5l216.9 125.7l216.6-125.7L256.5 87.9zM31 227.4v71l218 125.7v-71L31 227.4zm450 .2L265 353.1V424l216-125.5v-70.9z"/>`),
		g.Group(children),
	)
}