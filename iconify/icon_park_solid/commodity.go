package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commodity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCommodity0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M9.781 9.752A2 2 0 0 1 11.766 8h24.468a2 2 0 0 1 1.985 1.752l3.5 28A2 2 0 0 1 39.734 40H8.266a2 2 0 0 1-1.985-2.248l3.5-28Z"/><path stroke="#000" stroke-linecap="round" d="M15 18s2 4 9 4s9-4 9-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCommodity0)"/>`),
		g.Group(children),
	)
}