package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M177.9 494.1c-18.7 18.7-49.1 18.7-67.9 0l-92.1-92.2c-18.7-18.7-18.7-49.1 0-67.9l50.7-50.7l48 48c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-48-48l41.4-41.4l48 48c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-48-48l41.4-41.4l48 48c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-48-48l41.4-41.4l48 48c6.2 6.2 16.4 6.2 22.6 0s6.2-16.4 0-22.6l-48-48L333.9 18C352.6-.7 383-.7 401.8 18l92.1 92.1c18.7 18.7 18.7 49.1 0 67.9l-316 316.1z"/>`),
		g.Group(children),
	)
}