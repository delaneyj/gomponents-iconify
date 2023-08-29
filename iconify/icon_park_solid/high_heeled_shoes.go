package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighHeeledShoes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHighHeeledShoes0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M29.183 20C31.264 15.054 36.35 10.667 39 8c1.104.667 5 2.604 5 7c0 4-1.455 7.111-3 8l-5.987 4.191a23 23 0 0 0-6.533 7.01L25 40H4v-4c2.429-1.333 9.82-5.867 13-8c7 4 10.5-4 12.183-8Z"/><path d="M43 21v19"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHighHeeledShoes0)"/>`),
		g.Group(children),
	)
}