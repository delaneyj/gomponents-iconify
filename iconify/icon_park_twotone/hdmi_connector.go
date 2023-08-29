package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdmiConnector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHdmiConnector0"><g fill="none"><circle cx="24" cy="24" r="20" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 43a5 5 0 0 0-10 0"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 42.11A19.923 19.923 0 0 0 24 44c3.04 0 5.92-.678 8.5-1.89"/><circle cx="15" cy="15" r="3" fill="#fff"/><circle cx="11" cy="23" r="3" fill="#fff"/><circle cx="24" cy="11" r="3" fill="#fff"/><circle cx="33" cy="15" r="3" fill="#fff"/><circle cx="37" cy="23" r="3" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHdmiConnector0)"/>`),
		g.Group(children),
	)
}