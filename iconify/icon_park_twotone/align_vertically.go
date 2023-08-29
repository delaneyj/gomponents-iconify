package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignVertically0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M17 7h14v34H17z"/><path stroke-linecap="round" d="M42 24H6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignVertically0)"/>`),
		g.Group(children),
	)
}