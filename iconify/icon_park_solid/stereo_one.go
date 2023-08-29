package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSStereoOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="30" height="38" x="9" y="5" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 18h30"/><circle cx="24" cy="30" r="6" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSStereoOne0)"/>`),
		g.Group(children),
	)
}