package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chyrp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-446-647q-37 34-70 91t-51.5 107.5t-38.5 98t-34 62.5q-33 32-64.5 69t-48.5 60t-15 25q13 12 33.5 1t50.5-44.5t61.5-72.5t80.5-96t94-102q46-46 88-86.5t62-60.5t42-45q29-26 77-62t51-34q5 4-23 45.5t-91 117t-142 157.5q-95 100-192 192q37-9 69-25t63-40.5t54.5-46.5t60.5-58.5t64-60.5q53-49 100-112t76-118.5t50.5-104t29.5-77.5l9-29q-42 39-102.5 69t-112.5 48.5t-116.5 52.5t-114.5 79z"/>`),
		g.Group(children),
	)
}