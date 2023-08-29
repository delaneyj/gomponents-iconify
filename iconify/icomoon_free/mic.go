package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 11A2.5 2.5 0 0 0 10 8.5v-6a2.5 2.5 0 1 0-5 0v6A2.5 2.5 0 0 0 7.5 11zM11 7v1.5a3.5 3.5 0 1 1-7 0V7H3v1.5a4.5 4.5 0 0 0 4 4.472V15H5v1h5v-1H8v-2.028A4.5 4.5 0 0 0 12 8.5V7h-1z"/>`),
		g.Group(children),
	)
}