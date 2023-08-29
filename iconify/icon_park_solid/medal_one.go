package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMedalOne0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m6 6l10 12M42 6L32 18M31 4l-5 12M17 4l5 12"/><circle cx="24" cy="30" r="14" fill="#fff" stroke="#fff"/><circle cx="24" cy="30" r="7" fill="#000" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMedalOne0)"/>`),
		g.Group(children),
	)
}