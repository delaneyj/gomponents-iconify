package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 188h-28v-16h20a12 12 0 0 0 11.64-14.91l-32-128a12 12 0 0 0-23.28 0l-32 128A12 12 0 0 0 160 172h20v16h-64v-12a12 12 0 0 0 0-24v-12a12 12 0 0 0 0-24H32a12 12 0 0 0 0 24v12a12 12 0 0 0 0 24v12h-8a12 12 0 0 0 0 24h208a12 12 0 0 0 0-24ZM192 81.48L208.63 148h-33.26ZM56 140h36v12H56Zm0 36h36v12H56Zm60-76a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-40a8 8 0 1 1-8 8a8 8 0 0 1 8-8Z"/>`),
		g.Group(children),
	)
}