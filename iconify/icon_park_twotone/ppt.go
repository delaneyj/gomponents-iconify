package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ppt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPpt0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 8h40"/><path fill="#555" fill-rule="evenodd" d="M8 8h32v26H8V8Z" clip-rule="evenodd"/><path d="m22 16l5 5l-5 5m-6 16l8-8l8 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPpt0)"/>`),
		g.Group(children),
	)
}