package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinterestSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 1 1 4.584 6.912l1.911-4.367A2.58 2.58 0 0 0 8.5 11A3.5 3.5 0 0 0 12 7.5V7a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v.5c0 .896.337 1.715.891 2.333l.745-.666A2.489 2.489 0 0 1 4 7.5V7a3 3 0 0 1 3-3h1a3 3 0 0 1 3 3v.5A2.5 2.5 0 0 1 8.5 10A1.58 1.58 0 0 1 7 8.919l-.005-.016l.963-2.203l-.916-.4l-3.352 7.66A7.496 7.496 0 0 1 0 7.5Z"/>`),
		g.Group(children),
	)
}