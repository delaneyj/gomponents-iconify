package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multicast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMulticast0"><g fill="none" stroke="#fff"><path fill="#fff" stroke-width="4" d="M24 15a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-width="4" d="M24 15v8"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 33V23l23 .013V33"/><path stroke-linecap="round" stroke-width="5" d="M41 41v1M7 41v1m11-1v1m12-1v1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMulticast0)"/>`),
		g.Group(children),
	)
}