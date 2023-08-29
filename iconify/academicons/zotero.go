package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zotero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M31.264 64h276.752v54.624L79.072 401.392h228.944V448H11.984v-53.424l229.76-283.968H31.264Z"/>`),
		g.Group(children),
	)
}