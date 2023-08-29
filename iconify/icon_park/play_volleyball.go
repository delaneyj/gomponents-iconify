package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayVolleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-miterlimit="2" stroke-width="4" d="M29 14C31.7614 14 34 11.7614 34 9C34 6.23858 31.7614 4 29 4C26.2386 4 24 6.23858 24 9C24 11.7614 26.2386 14 29 14Z"/><path fill="#000" d="M41 24C42.6569 24 44 22.6569 44 21C44 19.3431 42.6569 18 41 18C39.3431 18 38 19.3431 38 21C38 22.6569 39.3431 24 41 24Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M16 24L13.73 35.26C13.52 36.57 12.12 37.31 10.92 36.76L4 34"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M37.9996 30.0002L23.7596 19.0703C22.9896 18.5003 21.9296 18.5603 21.2196 19.2003L16.3796 23.6303C15.9096 24.0603 15.9496 24.8003 16.4696 25.1803L25.9996 32.1503C26.7196 32.6803 27.0096 33.6303 26.6896 34.4703L23.1196 44.0003"/></g>`),
		g.Group(children),
	)
}