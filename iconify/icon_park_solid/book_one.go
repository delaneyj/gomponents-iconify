package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBookOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M7 37V11a6 6 0 0 1 6-6h22v26H13c-3.3 0-6 2.684-6 6Z"/><path stroke-linecap="round" d="M35 31H13a6 6 0 0 0 0 12h28V7M14 37h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBookOne0)"/>`),
		g.Group(children),
	)
}