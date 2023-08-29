package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalWatches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDigitalWatches0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="22" height="22" x="13" y="13" fill="#555" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 35v7a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-7m0-22V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v7m6 8h2m-18 3h2m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDigitalWatches0)"/>`),
		g.Group(children),
	)
}