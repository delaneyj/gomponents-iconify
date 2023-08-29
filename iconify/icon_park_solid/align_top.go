package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignTop0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M17 14.5h14v28H17z"/><path stroke-linecap="round" d="M42 6.5H6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignTop0)"/>`),
		g.Group(children),
	)
}