package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cancer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCancer0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="37" cy="17" r="6" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 13s6-8 16-8s16 6 16 6"/><circle cx="11" cy="31" r="6" fill="#fff" transform="rotate(-180 11 31)"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 35s-6 8-16 8s-16-6-16-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCancer0)"/>`),
		g.Group(children),
	)
}