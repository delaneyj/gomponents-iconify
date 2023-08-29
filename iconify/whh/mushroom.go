package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mushroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M897.5 511q-52.5 0-110-16T666 463t-132.5-16t-202 33T141 511q-197-6-117-128q30-46 64.5-89.5T179 194t114.5-96.5t136.5-69T587 0q101 0 184.5 30T910 112t85 122t30 149q0 62-37.5 95t-90 33zM534 511q49 0 100 9q-35 177 7 407q8 40-1.5 68t-36.5 28H396q-29 0-51.5-26.5T321 934q-6-213 58-401q93-22 155-22z"/>`),
		g.Group(children),
	)
}