package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrace0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6C14.059 6 6 14.059 6 24s8.059 18 18 18s18-8.059 18-18"/><path stroke-linecap="round" d="M24 15a9 9 0 1 0 9 9m-9 0l6.3-6.306"/><path fill="#fff" d="M30.3 11.426V17.7h6.325L42 12.3h-6.297V6L30.3 11.426Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrace0)"/>`),
		g.Group(children),
	)
}