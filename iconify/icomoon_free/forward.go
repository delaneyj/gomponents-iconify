package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.096 0C2.319 3.219 2.02 8.13 9 7.966V4l6 6l-6 6v-3.881C.641 12.337-.29 4.741 4.096 0z"/>`),
		g.Group(children),
	)
}