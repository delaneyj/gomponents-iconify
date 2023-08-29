package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHexagons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSThreeHexagons0"><g fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 27l-10-6l-10 6v12l10 6l10-6V27Zm20 0l-10-6l-10 6v12l10 6l10-6V27Z"/><path d="M34 9L24 3L14 9v12l10 6l10-6V9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSThreeHexagons0)"/>`),
		g.Group(children),
	)
}