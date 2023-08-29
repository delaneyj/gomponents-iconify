package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTProjectorTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M4 12a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 19h6m-6 6h4m-2 15v-6m20 6v-6"/><circle cx="31" cy="22" r="5" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTProjectorTwo0)"/>`),
		g.Group(children),
	)
}