package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GtaIii(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.427 17.356a2.079 2.079 0 0 0 2.073-2.073V10.3a2.079 2.079 0 0 0-2.073-2.073H6.573A2.079 2.079 0 0 0 4.5 10.3v4.983a2.079 2.079 0 0 0 2.073 2.073h1.113v13.487H6.573A2.079 2.079 0 0 0 4.5 32.916v4.983a2.079 2.079 0 0 0 2.073 2.073h34.854a2.079 2.079 0 0 0 2.073-2.073v-4.983a2.079 2.079 0 0 0-2.073-2.073h-1.113V17.356ZM15.44 30.843V17.356h4.683v13.487Zm17.12 0h-4.683V17.356h4.684Z"/>`),
		g.Group(children),
	)
}