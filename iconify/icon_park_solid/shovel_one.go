package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShovelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShovelOne0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M11 4h26l-2.833 20H13.833L11 4Z"/><path stroke="#000" d="M21 11v6m6-6v6"/><path stroke="#fff" d="M24 24v20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShovelOne0)"/>`),
		g.Group(children),
	)
}