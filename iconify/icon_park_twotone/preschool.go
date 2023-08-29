package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preschool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPreschool0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" d="M7 35h34a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2Z"/><path d="M14 14v14m20-14v14M24 17v8m-4-4h8"/><path stroke-linejoin="round" d="M4 41h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPreschool0)"/>`),
		g.Group(children),
	)
}