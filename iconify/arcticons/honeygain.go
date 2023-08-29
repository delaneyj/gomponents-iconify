package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Honeygain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.012 18.455a10.108 10.108 0 0 1 0 20.217c-5.388 0-16.96-1.92-20.079-7.759"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.012 18.455H10.729a6.23 6.23 0 1 0 3.874 11.072l.017.021Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.971 34.794l-8.264 2.859l4.226-6.74m13.8-7.873c-1.92 3.472-1.737 11.395 3.998 15.538M17.383 27.557c-.74 1.52.053 6.2 2.313 9.595"/><circle cx="31.962" cy="25.732" r="2.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="41.448" cy="15.309" r="2.052" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.985" cy="11.38" r="2.052" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.26 18.99c.985-1.223 3.242-3.457 6.112-3.79m-8.795 3.271a8.934 8.934 0 0 1 2.733-5.932"/>`),
		g.Group(children),
	)
}