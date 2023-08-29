package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globalfreecall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M22.433 18.551v12.337A4.112 4.112 0 0 1 18.321 35h0a4.1 4.1 0 0 1-2.908-1.205"/><rect width="8.224" height="10.897" x="14.209" y="18.551" rx="4.112" ry="4.112" transform="rotate(180 18.321 24)"/><path d="M28.005 29.449v-13.57A2.879 2.879 0 0 1 30.883 13h0c1.417 0 2.301.42 2.908 1.204m-8.224 4.347h5.757"/></g>`),
		g.Group(children),
	)
}