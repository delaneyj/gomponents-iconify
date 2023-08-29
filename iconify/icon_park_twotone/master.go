package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Master(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMaster0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 38V6a2 2 0 0 0-2-2H12a2 2 0 0 0-2 2v32"/><rect width="28" height="12" x="10" y="32" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="6"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m20 16l4-4l4 4l4-4l-3 11H19l-3-11l4 4Z"/><circle cx="32" cy="38" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMaster0)"/>`),
		g.Group(children),
	)
}