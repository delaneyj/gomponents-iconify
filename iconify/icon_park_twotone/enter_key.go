package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEnterKey0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 44V4H24v16H4v24h40Z"/><path d="m21 28l-4 4l4 4"/><path d="M34 23v9H17"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEnterKey0)"/>`),
		g.Group(children),
	)
}