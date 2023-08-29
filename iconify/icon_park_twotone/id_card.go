package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIdCard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M42 8H6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Z"/><path fill="#555" d="M36 16h-8v8h8v-8Z"/><path stroke-linecap="round" d="M12 32h24M12 16h6m-6 8h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIdCard0)"/>`),
		g.Group(children),
	)
}