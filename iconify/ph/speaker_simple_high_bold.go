package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSimpleHighBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M157.27 21.22a12 12 0 0 0-12.64 1.31L75.88 76H32a20 20 0 0 0-20 20v64a20 20 0 0 0 20 20h43.88l68.75 53.47A12 12 0 0 0 164 224V32a12 12 0 0 0-6.73-10.78ZM140 199.47l-52.63-40.94A12 12 0 0 0 80 156H36v-56h44a12 12 0 0 0 7.37-2.53L140 56.54ZM204 104v48a12 12 0 0 1-24 0v-48a12 12 0 0 1 24 0Zm36-16v80a12 12 0 0 1-24 0V88a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}