package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlocksAndArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlocksAndArrows0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M20 6H6v14h14V6Zm0 22H6v14h14V28ZM42 6H28v14h14V6Z"/><path d="m28 28l14 14M28 28h14h-14Zm0 0v14v-14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlocksAndArrows0)"/>`),
		g.Group(children),
	)
}