package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m96 373.3l-96 96L42.7 512l96-96l74.7 74.7v-192h-192L96 373.3zm394.7-74.6h-192v192l74.7-74.7l96 96l42.7-42.7l-96-96l74.6-74.6zM42.7 0L0 42.7l96 96l-74.7 74.7h192v-192L138.7 96l-96-96zM416 138.7l96-96L469.3 0l-96 96l-74.7-74.7v192h192L416 138.7z"/>`),
		g.Group(children),
	)
}