package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gastromapalukehejlka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24c0 11.874-9.626 21.5-21.5 21.5S2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24ZM12 12l24 24m-24 0l24-24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.273 9.866c0-1.312-1.136-2.449-2.448-2.361c-1.224.087-2.186 1.224-2.186 2.448v2.186c0 1.312 1.05 2.361 2.361 2.361s2.36-1.05 2.36-2.36H24m10.681 8.36v7m4.638-7v7M34.681 24h4.638M20.5 40.5v-7l3.5 7l3.5-7v7m-18.25-20v7h3.5"/>`),
		g.Group(children),
	)
}