package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungGoodGuardians(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="33.25" cy="14.75" r="9.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24h18.5v18.5H24zm0 18.5C13.783 42.5 5.5 34.217 5.5 24H24v18.5ZM5.5 24C5.5 13.783 13.783 5.5 24 5.5V24H5.5Z"/>`),
		g.Group(children),
	)
}