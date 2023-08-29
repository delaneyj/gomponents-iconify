package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 64v43H0V64h256zm0 85v43H0v-43h256zM0 277v-42h171v42H0zM299 64h106v43h-64v192q0 26-18.5 45t-45 19t-45.5-19t-19-45.5t19-45t45-18.5q11 0 22 4V64z"/>`),
		g.Group(children),
	)
}