package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1v1h6V1h2v2.162l-1 3V15h1v3h2v5H5v-5h2v-3h1V6.162l-1-3V1h2Zm1 6v8h4V7h-4Zm4.28-2l.333-1H9.387l.334 1h4.558ZM9 17v1h6v-1H9Zm8 3H7v1h10v-1Z"/>`),
		g.Group(children),
	)
}