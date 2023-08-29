package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindChime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M55.177 38.627a19.673 19.673 0 0 0 .494-9.215A19.997 19.997 0 0 0 16 33a19.742 19.742 0 0 0 .824 5.63a2.005 2.005 0 0 0 1.924 1.43h34.504a2.007 2.007 0 0 0 1.925-1.434Z"/><path fill="#f1b31c" d="M39 41a3 3 0 0 1-6 0"/><path fill="#d0cfce" d="M16 54a14.423 14.423 0 0 0 14-4l4 6a14.423 14.423 0 0 1-14 4Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M55.177 38.627a19.673 19.673 0 0 0 .494-9.215A19.997 19.997 0 0 0 16 33a19.742 19.742 0 0 0 .824 5.63a2.005 2.005 0 0 0 1.924 1.43h34.504a2.007 2.007 0 0 0 1.925-1.434Z"/><path stroke-linecap="round" stroke-miterlimit="10" d="M36 13s.789-4.106-1-5"/><path stroke-miterlimit="10" d="M36 44s1 6-4 9"/><path stroke-miterlimit="10" d="M39 41a3 3 0 0 1-6 0"/><path stroke-linejoin="round" d="M16 54a14.423 14.423 0 0 0 14-4l4 6a14.423 14.423 0 0 1-14 4Z"/></g>`),
		g.Group(children),
	)
}