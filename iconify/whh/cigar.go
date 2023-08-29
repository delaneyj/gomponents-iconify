package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cigar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1011.27 955q-7 6-18 17.5t-20.5 20.5t-17.5 18q-12 12-46 13t-62-14q-241-154-491-411l36-36q28 13 56 13q53 0 90.5-37.5t37.5-90.5q0-28-13-56l36-36q257 249 411 491q15 28 14 62t-13 46zm-563-635q-53 0-90.5 37.5t-37.5 90.5q0 28 13 56l-35 35q-174-189-284-362q-15-28-14-62t13-46l56-56q12-12 46-13t62 14q174 111 362 284l-35 35q-28-13-56-13z"/>`),
		g.Group(children),
	)
}