package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridBevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 2V1H1v13h1v1h13V2h-1zM5 13H2v-3h3v3zm0-4H2V6h3v3zm0-4H2V2h3v3zm4 8H6v-3h3v3zm0-4H6V6h3v3zm0-4H6V2h3v3zm4 8h-3v-3h3v3zm0-4h-3V6h3v3zm0-4h-3V2h3v3z"/>`),
		g.Group(children),
	)
}