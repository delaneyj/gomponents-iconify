package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPowerpoint0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M19 15h10v10H19z"/><path stroke-linecap="round" d="M19 33V15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPowerpoint0)"/>`),
		g.Group(children),
	)
}