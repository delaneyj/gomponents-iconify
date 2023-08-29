package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 8C18.1 8 17 9.1 17 10.5c0 1.2.9 2.2 2 2.4v2.6c0 3-2.5 5.5-5.5 5.5S8 18.5 8 15.5V15c2.8-.3 5-2.6 5-5.5v-7c0-.3-.2-.5-.5-.5h-2c-.3 0-.5.2-.5.5s.2.5.5.5H12v6.5C12 12 10 14 7.5 14S3 12 3 9.5V3h1.5c.3 0 .5-.2.5-.5S4.8 2 4.5 2h-2c-.3 0-.5.2-.5.5v7c0 2.9 2.2 5.2 5 5.5v.5c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-2.6c1.1-.2 2-1.2 2-2.4C22 9.1 20.9 8 19.5 8zm0 4c-.8 0-1.5-.7-1.5-1.5S18.7 9 19.5 9s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5z"/>`),
		g.Group(children),
	)
}