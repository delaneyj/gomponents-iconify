package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCopyright0"><g fill="none"><circle cx="24" cy="24" r="20" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#555" d="M24 17h-4v7h4c3 0 4-2 4-3.5S27 17 24 17Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 31v-7m0 0v-7h4c3 0 4 2 4 3.5S27 24 24 24h-1m-3 0h3m5 7l-5-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCopyright0)"/>`),
		g.Group(children),
	)
}