package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShake0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="m4 10l4 4.667l-4 4.666L8 24l-4 4.667l4 4.666L4 38m40-28l-4 4.667l4 4.666L40 24l4 4.667l-4 4.666L44 38"/><path fill="#fff" stroke="#fff" d="M34 6H14v36h20V6Z"/><path stroke="#000" stroke-linecap="round" d="M22 35h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShake0)"/>`),
		g.Group(children),
	)
}