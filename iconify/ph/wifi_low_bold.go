package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLowBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 204a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm31.06-48.7a80 80 0 0 0-94.12 0a12 12 0 1 0 14.13 19.4a56 56 0 0 1 65.86 0a12 12 0 1 0 14.13-19.4Z"/>`),
		g.Group(children),
	)
}