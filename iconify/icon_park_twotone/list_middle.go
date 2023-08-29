package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTListMiddle0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M8 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 28a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-linecap="round" d="M20 24h24M20 38h24M20 10h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTListMiddle0)"/>`),
		g.Group(children),
	)
}