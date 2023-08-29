package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAppSwitch0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M34 4H14v40h20V4Z"/><path d="M44 8H34v32h10V8ZM14 8H4v32h10V8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAppSwitch0)"/>`),
		g.Group(children),
	)
}