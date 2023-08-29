package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Float(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFloat0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M24 40c11.046 0 20-7.163 20-16S35.046 8 24 8S4 15.163 4 24s8.954 16 20 16Z"/><path fill="#555" d="M24 28c5.523 0 10-2.686 10-6s-4.477-6-10-6s-10 2.686-10 6s4.477 6 10 6Z"/><path stroke-linecap="round" d="M24 16V8m8 10s2.625-4 7-4m-23 4s-2-4-7-4m9 13s-6 2-7 9m19-9s6.5 2 7 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFloat0)"/>`),
		g.Group(children),
	)
}