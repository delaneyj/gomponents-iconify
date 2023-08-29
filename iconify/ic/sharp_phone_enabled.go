package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPhoneEnabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 15.46l5.27-.61l2.52 2.52c2.83-1.44 5.15-3.75 6.59-6.59l-2.53-2.53l.61-5.25h5.51C21.55 13.18 13.18 21.55 3 20.97v-5.51z"/>`),
		g.Group(children),
	)
}