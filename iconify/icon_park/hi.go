package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 43H17.476C17.3873 43 17.3049 42.9541 17.2581 42.8788L7.86011 27.7273C6.79115 26.0039 7.14475 23.7577 8.69148 22.446C10.6306 20.8016 13.584 21.3036 14.871 23.4963L17.3333 27.6914V8.57577C17.3333 6.71037 19.0177 5.29724 20.8547 5.62142L37.5214 8.56259C38.9549 8.81558 40 10.0612 40 11.5169V17.8148V40C40 41.6569 38.6569 43 37 43Z"/>`),
		g.Group(children),
	)
}