package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 195q0 36-19 70l-26-27q9-21 9-43h36zm-85 3L128 71v-4q0-27 18.5-45.5T192 3t45.5 18.5T256 67v131zM27 24l357 357l-27 27l-89-89q-26 15-55 19v70h-42v-70q-54-8-91-49t-37-94h36q0 46 33.5 77t79.5 31q25 0 49-11l-35-35q-7 2-14 2q-27 0-45.5-19T128 195v-16L0 51z"/>`),
		g.Group(children),
	)
}