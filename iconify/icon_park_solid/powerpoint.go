package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPowerpoint0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M19 15h10v10H19z"/><path stroke="#000" stroke-linecap="round" d="M19 33V15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPowerpoint0)"/>`),
		g.Group(children),
	)
}