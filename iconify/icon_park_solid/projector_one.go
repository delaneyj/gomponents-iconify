package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSProjectorOne0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 21V10H10v11"/><rect width="40" height="16" x="4" y="21" fill="#fff" stroke="#fff" stroke-width="4" rx="2"/><rect width="4" height="4" x="14" y="27" fill="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 29h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSProjectorOne0)"/>`),
		g.Group(children),
	)
}