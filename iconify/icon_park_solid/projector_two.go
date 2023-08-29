package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSProjectorTwo0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 12a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V12Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 19h6m-6 6h4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 40v-6m20 6v-6"/><circle cx="31" cy="22" r="5" fill="#000" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSProjectorTwo0)"/>`),
		g.Group(children),
	)
}