package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 8V4.9C12 2.7 10.4 1 8.2 1h-.3C5.8 1 4 2.7 4 4.9V8H3l.1 5S3 16 8 16s5-3 5-3V8h-1zm-3 6H8v-2c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1v3zm1-6H6V4.9C6 3.8 6.9 3 7.9 3h.3c1 0 1.8.8 1.8 1.9V8z"/>`),
		g.Group(children),
	)
}