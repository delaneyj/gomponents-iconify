package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelephoneEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M3 9.4a.73.73 0 0 0 1 0l.74-.74a.73.73 0 0 0 0-1zm4.56-4.6a.73.73 0 0 0 1 0l.71-.71a.73.73 0 0 0 0-1zM5.88 3.57L3.57 5.88a.37.37 0 0 0 0 .52l.43.44L2.26 8.6a2.27 2.27 0 0 1-.73-1.34v-1a1.345 1.345 0 0 1 .52-1l3.21-3.21a1.345 1.345 0 0 1 1-.52h1a2.27 2.27 0 0 1 1.34.73L6.84 4l-.44-.43a.37.37 0 0 0-.52 0" fill="currentColor"/>`),
		g.Group(children),
	)
}