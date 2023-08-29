package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 0v16h8V0H5zm1 2h2v1H6V2zm0 6h2v1H6V8zm2 7H6v-1h2v1zm1-3H6v-1h3v1zm0-6H6V5h3v1z"/>`),
		g.Group(children),
	)
}