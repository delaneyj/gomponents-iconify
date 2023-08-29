package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandleDown0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linecap="round" rx="3"/><path d="M34 20L24 30L14 20h20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandleDown0)"/>`),
		g.Group(children),
	)
}