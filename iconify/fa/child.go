package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Child(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1124 388L832 680v824q0 46-33 79t-79 33t-79-33t-33-79v-384h-64v384q0 46-33 79t-79 33t-79-33t-33-79V680L28 388Q0 360 0 320t28-68q29-28 68.5-28t67.5 28l228 228h368l228-228q28-28 68-28t68 28q28 29 28 68.5t-28 67.5zM800 224q0 93-65.5 158.5T576 448t-158.5-65.5T352 224t65.5-158.5T576 0t158.5 65.5T800 224z"/>`),
		g.Group(children),
	)
}