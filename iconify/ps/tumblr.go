package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 378q-12-7-17-20q-3-7-3-49V190h108v-82H172V2h-66q-5 37-16 58q-12 24-31 40q-14 12-53 25v65h63v162q0 33 7 48q6 17 24 32t42 22q21 8 50 8q32 0 52-5q28-7 54-20v-72q-34 22-70 22q-20 0-36-9z"/>`),
		g.Group(children),
	)
}