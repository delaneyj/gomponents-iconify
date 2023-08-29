package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m220 88.59l-176.08-54A14 14 0 0 0 26 48v144a14 14 0 0 0 14 14a14.59 14.59 0 0 0 4-.59L130 179v13a14 14 0 0 0 14 14h32a14 14 0 0 0 14-14v-31.38l30-9.19A14.07 14.07 0 0 0 230 138v-36a14 14 0 0 0-10-13.41ZM40.56 193.92a2 2 0 0 1-1.76-.32a2 2 0 0 1-.8-1.6V48a2 2 0 0 1 .8-1.6A2.05 2.05 0 0 1 40 46a1.79 1.79 0 0 1 .49.08L130 73.53v92.95ZM178 192a2 2 0 0 1-2 2h-32a2 2 0 0 1-2-2v-16.65l36-11Zm40-54a2 2 0 0 1-1.44 1.92h-.08L142 162.8V77.21l74.56 22.87A2 2 0 0 1 218 102Z"/>`),
		g.Group(children),
	)
}