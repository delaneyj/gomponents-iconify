package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTViewList0"><g fill="none"><rect width="32" height="40" x="8" y="4" fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 14h12M21 24h12M21 34h12"/><path fill="#fff" fill-rule="evenodd" d="M15 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTViewList0)"/>`),
		g.Group(children),
	)
}