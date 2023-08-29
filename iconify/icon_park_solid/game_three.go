package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGameThree0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 30v3a7 7 0 0 1-7 7v0a7 7 0 0 1-7-7V19m24 11v3a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7V19"/><rect width="38" height="22" x="5" y="8" fill="#fff" stroke="#fff" stroke-width="4" rx="11"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 19h-8m4-4v8"/><rect width="4" height="4" x="32" y="15" fill="#000" rx="2"/><rect width="4" height="4" x="28" y="20" fill="#000" rx="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGameThree0)"/>`),
		g.Group(children),
	)
}