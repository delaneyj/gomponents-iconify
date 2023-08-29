package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreastPump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBreastPump0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M35 25c0-4-6-6-6-6v-5H17v5s-6 2-6 6v19h24V25Z"/><path d="m20 4l-7 6m10 4l-6-7m9 1h9v7l6 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBreastPump0)"/>`),
		g.Group(children),
	)
}