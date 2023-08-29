package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLabel0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 4h30v40L24 33.429L9 44V4Z"/><path fill="#555" d="M9 4h30v12H9z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLabel0)"/>`),
		g.Group(children),
	)
}