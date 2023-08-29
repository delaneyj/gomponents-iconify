package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 52v40a12 12 0 0 1-24 0V64h-28a12 12 0 0 1 0-24h40a12 12 0 0 1 12 12ZM76 192H48v-28a12 12 0 0 0-24 0v40a12 12 0 0 0 12 12h40a12 12 0 0 0 0-24Zm144-40a12 12 0 0 0-12 12v28h-28a12 12 0 0 0 0 24h40a12 12 0 0 0 12-12v-40a12 12 0 0 0-12-12ZM36 104a12 12 0 0 0 12-12V64h28a12 12 0 0 0 0-24H36a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12Zm52-24a12 12 0 0 0-12 12v72a12 12 0 0 0 24 0V92a12 12 0 0 0-12-12Zm92 84V92a12 12 0 0 0-24 0v72a12 12 0 0 0 24 0Zm-52-84a12 12 0 0 0-12 12v72a12 12 0 0 0 24 0V92a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}