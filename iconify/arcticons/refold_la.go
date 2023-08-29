package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefoldLa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M37.117 18.284L33.16 31.623l-16.63 4.452l-.713-.705l-.951-23.445l19.627 15.207M15.818 35.37L3.5 23.166l11.972 3.671m12.451-4.796l3.189-3.265"/><path d="m44.5 17.678l-8.537-3.869l-4.851 4.967L44.5 17.678Z"/></g>`),
		g.Group(children),
	)
}