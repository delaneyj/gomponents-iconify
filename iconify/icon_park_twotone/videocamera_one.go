package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVideocameraOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="31" height="20" x="4" y="21" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 27h9v8h-9z"/><circle cx="27" cy="13" r="5" fill="#555" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13" cy="13" r="5" fill="#555" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m35 35l9 4V23l-9 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVideocameraOne0)"/>`),
		g.Group(children),
	)
}