package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grooveshark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2zm124 301q-8 0-17-9t-13-18l-5-9q-1-4-3.5-10t-13-22t-25-29.5T238 179t-59-18q2 33 .5 58t-6 39.5t-11 25t-14 13.5t-15 5t-14 1t-11.5 0q-16 1-25.5-19.5T73 232q0-66 46.5-112.5T232 73t111.5 46.5T391 232q1 16-11 43.5T356 303z"/>`),
		g.Group(children),
	)
}