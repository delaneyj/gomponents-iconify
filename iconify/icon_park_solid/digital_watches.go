package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalWatches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDigitalWatches0"><g fill="none" stroke-width="4"><rect width="22" height="22" x="13" y="13" fill="#fff" stroke="#fff" rx="3"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M29 35v7a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-7m0-22V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7m6 8h2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M19 24h2m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDigitalWatches0)"/>`),
		g.Group(children),
	)
}