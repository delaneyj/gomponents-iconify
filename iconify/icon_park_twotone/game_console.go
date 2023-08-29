package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGameConsole0"><g fill="none"><rect width="28" height="40" x="10" y="4" stroke="#fff" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 34h8m-4-4v8"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M16 10h16v9H16z"/><circle cx="31" cy="30" r="2" fill="#fff"/><circle cx="31" cy="38" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGameConsole0)"/>`),
		g.Group(children),
	)
}