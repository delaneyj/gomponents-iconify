package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeEvergreenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M233.47 184.63L192.54 132H208a12 12 0 0 0 9.51-19.32l-80-104a12 12 0 0 0-19 0l-80 104A12 12 0 0 0 48 132h15.46l-40.93 52.63A12 12 0 0 0 32 204h84v36a12 12 0 0 0 24 0v-36h84a12 12 0 0 0 9.47-19.37ZM56.54 180l40.93-52.63A12 12 0 0 0 88 108H72.37L128 35.68L183.63 108H168a12 12 0 0 0-9.47 19.37L199.46 180Z"/>`),
		g.Group(children),
	)
}