package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFork0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M37 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-26 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm13 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M11 12v3c0 7 13 10 13 17v4v-4c0-7 13-10 13-17v-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFork0)"/>`),
		g.Group(children),
	)
}