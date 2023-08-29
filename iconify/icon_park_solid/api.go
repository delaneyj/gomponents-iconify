package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Api(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSApi0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m37 22l-3 3l-11-11l3-3c1.5-1.5 7-4 11 0s1.5 9.5 0 11Z"/><path d="m42 6l-5 5"/><path fill="#fff" d="m11 26l3-3l11 11l-3 3c-1.5 1.5-7 4-11 0s-1.5-9.5 0-11Z"/><path d="m23 32l4-4M6 42l5-5m5-12l4-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSApi0)"/>`),
		g.Group(children),
	)
}