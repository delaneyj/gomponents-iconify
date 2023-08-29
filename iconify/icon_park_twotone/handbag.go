package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handbag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandbag0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 14a6 6 0 0 1 6-6h9a6 6 0 0 1 6 6v2H14v-2Z"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m4 25l19.515 4.879c.318.08.652.08.97 0L44 25v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V25Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 27v-9a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v9"/><path fill="#fff" d="M26.5 23a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandbag0)"/>`),
		g.Group(children),
	)
}