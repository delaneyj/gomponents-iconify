package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMedalOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m6 6l10 12M42 6L32 18M31 4l-5 12M17 4l5 12"/><circle cx="24" cy="30" r="14" fill="#555"/><circle cx="24" cy="30" r="7" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMedalOne0)"/>`),
		g.Group(children),
	)
}