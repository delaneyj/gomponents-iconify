package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTool0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 16c0 6.627-5.373 12-12 12c-2.027 0-3.936-.503-5.61-1.39L9 44l-5-5l17.39-17.39A11.948 11.948 0 0 1 20 16c0-6.627 5.373-12 12-12c2.027 0 3.936.502 5.61 1.39L30 13l5 5l7.61-7.61A11.948 11.948 0 0 1 44 16Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTool0)"/>`),
		g.Group(children),
	)
}