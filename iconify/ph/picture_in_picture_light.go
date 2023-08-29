package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 50H40a14 14 0 0 0-14 14v128a14 14 0 0 0 14 14h176a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14ZM38 192V64a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2v58h-74a14 14 0 0 0-14 14v58H40a2 2 0 0 1-2-2Zm178 2h-74v-58a2 2 0 0 1 2-2h74v58a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}