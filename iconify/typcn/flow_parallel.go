package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowParallel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16.184V7.816A2.997 2.997 0 0 0 20 5c0-1.654-1.346-3-3-3s-3 1.346-3 3c0 1.302.839 2.401 2 2.815v8.369A2.997 2.997 0 0 0 14 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816zM17 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1zm0 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2zM10 5c0-1.654-1.346-3-3-3S4 3.346 4 5c0 1.302.839 2.401 2 2.815v8.369A2.997 2.997 0 0 0 4 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816V7.816A2.997 2.997 0 0 0 10 5zM7 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1zm0 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}