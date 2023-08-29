package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="m4 0l-.34.34C3.55.45 1 3.03 1 5.22c0 1.65 1.35 3 3 3s3-1.35 3-3C7 3.04 4.45.45 4.34.34L4 0zM2.5 4.72c.28 0 .5.22.5.5c0 .55.45 1 1 1c.28 0 .5.22.5.5s-.22.5-.5.5c-1.1 0-2-.9-2-2c0-.28.22-.5.5-.5z"/>`),
		g.Group(children),
	)
}