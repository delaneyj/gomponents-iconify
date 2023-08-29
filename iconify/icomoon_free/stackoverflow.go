package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 10v6H0v-6h2v4h12v-4zM3 11h10v2H3zm.237-2.165l.433-1.953l9.763 2.164L13 10.999zM4.37 4.821l.845-1.813l9.063 4.226l-.845 1.813zm11.126.827l-1.218 1.587l-7.934-6.088L7.224 0h.91z"/>`),
		g.Group(children),
	)
}