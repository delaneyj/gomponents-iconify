package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWebPage0"><g fill="none"><rect width="40" height="32" x="4" y="8" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="#555" stroke="#fff" stroke-width="4" d="M4 11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 16 14)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWebPage0)"/>`),
		g.Group(children),
	)
}