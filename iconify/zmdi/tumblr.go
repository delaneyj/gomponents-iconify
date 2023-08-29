package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 177v-60q25-8 43-23.5T72 57T87 3h61v108h102v66H148v110q0 37 4 47.5t15 16.5q14 9 33 9q32 0 65-21v67q-28 13-50.5 18t-48.5 5q-29 0-51.5-7T76 401t-22.5-29.5T47 327V178H0v-1z"/>`),
		g.Group(children),
	)
}