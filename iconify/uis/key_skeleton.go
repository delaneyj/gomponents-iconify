package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeleton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 4.4l.7-.7c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L9.8 12.8C9 12.3 8 12 7 12c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5c0-1-.3-2-.8-2.8l5.6-5.6l2.1 2.1c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-2.1-2.1l1.4-1.4l.7.7c.4.4 1 .4 1.4 0s.4-1 0-1.4l-.7-.7zM7 20c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3z"/>`),
		g.Group(children),
	)
}