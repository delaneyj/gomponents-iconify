package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmchairBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 78.53V72a44.05 44.05 0 0 0-44-44H80a44.05 44.05 0 0 0-44 44v6.53a52 52 0 0 0 0 99V200a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20v-22.49a52 52 0 0 0 0-99ZM80 52h96a20 20 0 0 1 20 20v4.62A52.09 52.09 0 0 0 152.17 124h-48.34A52.09 52.09 0 0 0 60 76.62V72a20 20 0 0 1 20-20Zm126.81 103.86A12 12 0 0 0 196 167.8V196H60v-28.2a12 12 0 0 0-10.81-11.94A28 28 0 1 1 80 128v36a12 12 0 0 0 24 0v-16h48v16a12 12 0 0 0 24 0v-36a28 28 0 1 1 30.81 27.86Z"/>`),
		g.Group(children),
	)
}