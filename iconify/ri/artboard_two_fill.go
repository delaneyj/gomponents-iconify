package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtboardTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6h12v12H6V6Zm0-4h2v3H6V2Zm0 17h2v3H6v-3ZM2 6h3v2H2V6Zm0 10h3v2H2v-2ZM19 6h3v2h-3V6Zm0 10h3v2h-3v-2ZM16 2h2v3h-2V2Zm0 17h2v3h-2v-3Z"/>`),
		g.Group(children),
	)
}