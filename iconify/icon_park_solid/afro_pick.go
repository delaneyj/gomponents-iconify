package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfroPick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAfroPick0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21 4.372L4.03 21.342L26.656 43.97L43.627 27m-16.97-16.971l-16.97 16.97m22.626-11.314l-16.97 16.971M37.97 21.343L21 38.313"/><path fill="#fff" d="m16.05 41.849l4.243-4.243l-9.9-9.9l-4.242 4.243l2.121 2.122l1.414 4.242l4.243 1.414l2.121 2.122Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAfroPick0)"/>`),
		g.Group(children),
	)
}