package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 8h6v2H5v6H3v-6H1zm14-4h-3.934v12H8.933V4H4.999V2h10z"/>`),
		g.Group(children),
	)
}