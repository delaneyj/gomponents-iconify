package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zencom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.6 17.852A21.498 21.498 0 1 1 26.767 2.68"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.593 17.821c-3.173-12.197-19.646-.25-6.973 5.729"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.622 23.556a13.405 13.405 0 0 1-12.346 13.456c-7.168.733-13.724-3.904-15.062-10.653a13.162 13.162 0 0 1 9.947-15.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.71 11.31s.813.017.952-1.63c.353-4.486 2-3.448 6.346-3.513c2.656-.374 2.334-3.237-.216-3.48"/><ellipse cx="33.205" cy="6.446" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.266" ry="2.243"/>`),
		g.Group(children),
	)
}