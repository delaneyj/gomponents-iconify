package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPlug0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 14h36v10c-4 8-10 12-18 12S10 32 6 24V14Z"/><path stroke="#fff" stroke-linecap="round" d="m33 34l-1 10H16l-1-10"/><path stroke="#000" stroke-linecap="round" d="M22 24h4"/><path stroke="#fff" stroke-linecap="round" d="M16 4v8m16-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPlug0)"/>`),
		g.Group(children),
	)
}