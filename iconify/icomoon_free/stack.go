package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 5L8 1L0 5l8 4l8-4zM8 2.328L13.345 5L8 7.672L2.655 5L8 2.328zm6.398 4.871L16 8l-8 4l-8-4l1.602-.801L8 10.398zm0 3L16 11l-8 4l-8-4l1.602-.801L8 13.398z"/>`),
		g.Group(children),
	)
}