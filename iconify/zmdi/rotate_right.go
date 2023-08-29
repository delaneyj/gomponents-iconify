package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m246 102l-97 95v-83q-45 8-75.5 43.5T43 240t30.5 82.5T149 366v43q-63-8-106-56T0 240t43-113t106-56V5zm94 117h-43q-5-29-22-53l30-30q29 37 35 83zM192 366q28-5 52-22l31 31q-37 28-83 34v-43zm83-52q17-24 22-53h43q-6 46-35 83z"/>`),
		g.Group(children),
	)
}