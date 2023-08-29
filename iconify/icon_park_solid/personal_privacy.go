package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalPrivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPersonalPrivacy0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#fff"/><path d="M4 41c0-8.837 8.059-16 18-16"/><path fill="#fff" d="M27 31h14v10H27z"/><path d="M37 31v-3a3 3 0 1 0-6 0v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPersonalPrivacy0)"/>`),
		g.Group(children),
	)
}