package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.677 2.603a1.75 1.75 0 0 1 2.644 0l8.246 9.504c.983 1.133.178 2.897-1.322 2.897H17v5.247A1.75 1.75 0 0 1 15.25 22h-6.5A1.75 1.75 0 0 1 7 20.25v-5.247H3.754c-1.5 0-2.305-1.764-1.322-2.897l8.245-9.504Zm1.511.983a.25.25 0 0 0-.378 0L3.566 13.09a.25.25 0 0 0 .189.414H7.75a.75.75 0 0 1 .75.75v5.997c0 .138.112.25.25.25h6.5a.25.25 0 0 0 .25-.25v-5.997a.75.75 0 0 1 .75-.75h3.995a.25.25 0 0 0 .188-.414l-8.245-9.504Z"/>`),
		g.Group(children),
	)
}