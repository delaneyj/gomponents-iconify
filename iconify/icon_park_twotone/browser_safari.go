package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserSafari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBrowserSafari0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m15 33l4.5-13.5L33 15l-4.5 13.5L15 33Z"/><path fill="#fff" fill-rule="evenodd" d="M24 26a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBrowserSafari0)"/>`),
		g.Group(children),
	)
}