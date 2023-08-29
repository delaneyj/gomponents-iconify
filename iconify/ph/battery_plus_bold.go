package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryPlusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 128a12 12 0 0 1-12 12h-12v12a12 12 0 0 1-24 0v-12H92a12 12 0 0 1 0-24h12v-12a12 12 0 0 1 24 0v12h12a12 12 0 0 1 12 12Zm72-48v96a28 28 0 0 1-28 28H28a28 28 0 0 1-28-28V80a28 28 0 0 1 28-28h168a28 28 0 0 1 28 28Zm-24 0a4 4 0 0 0-4-4H28a4 4 0 0 0-4 4v96a4 4 0 0 0 4 4h168a4 4 0 0 0 4-4Zm44 12a12 12 0 0 0-12 12v48a12 12 0 0 0 24 0v-48a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}