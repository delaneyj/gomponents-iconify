package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cafeteria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 1v7c0 1.657-1.347 3-2.997 3H5.997A3.002 3.002 0 0 1 3 8V1m3 6V1m3 6V1M6 11v10.504C6 22.33 6.666 23 7.5 23c.828 0 1.5-.68 1.5-1.496V11m6 7v3.5a1.5 1.5 0 1 0 3-.005V15s3 0 3-3V7c0-3-2-5-6-5v16Z"/>`),
		g.Group(children),
	)
}