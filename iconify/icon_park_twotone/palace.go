package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPalace0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 18h40L24 6L4 18Z"/><path d="M44 42H4m5-24v24m10-24v24m10-24v24m10-24v24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPalace0)"/>`),
		g.Group(children),
	)
}