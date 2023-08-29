package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCaution0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M24 5L2 43h44L24 5Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M24 35v1m0-17l.008 10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCaution0)"/>`),
		g.Group(children),
	)
}