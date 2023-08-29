package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Africa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m201.56 19.495l-87.79 9.131l-73.745 94.814v52.676l56.186 61.805l64.615-13.344l49.164 9.832l-10.535 37.926l33.711 61.103l-16.855 42.842l39.79 116.225l53.62-8.768l49.164-55.484l4.213-38.629l31.605-23.879l-6.322-69.531l83.594-106.994l-51.989 7.263l-79.363-138.359l-125.016-8.428l-14.046-30.2zm252.346 319.8l-14.402 20.86l-13.408.496c-11.849 24.321-12.598 38.019-13.907 66.547l17.383 4.471l21.852-52.147l2.482-40.226z"/>`),
		g.Group(children),
	)
}