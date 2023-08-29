package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplytranslate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.901 17.989l3.165-9.507m3.033 9.535l-3.033-9.535m2.018 6.346h-4.131m-3.247 15.53h10.626m-5.313-2.105v2.105m2.974 0c0 2.908-3.375 7.72-6.75 8.455"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.777 32.998c.335 2.006 3.81 5.347 6.9 5.815"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.786 33.286a10.214 10.214 0 0 0 20.426.16l.002-.156a8.617 8.617 0 0 0-6.331-8.18l-7.79-2.227a8.538 8.538 0 0 1-6.287-7.53q-.025-.396-.019-.8a10.214 10.214 0 0 1 20.426.16"/>`),
		g.Group(children),
	)
}