package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBrowser0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 18v22a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V18"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M6 8a2 2 0 0 1 2-2h32a2 2 0 0 1 2 2v10H6V8Z"/><path fill="#fff" fill-rule="evenodd" d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBrowser0)"/>`),
		g.Group(children),
	)
}