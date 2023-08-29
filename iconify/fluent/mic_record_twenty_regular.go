package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicRecordTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 13c.07 0 .14-.002.21-.007c.11-.387.26-.757.448-1.104A2 2 0 0 1 7 10V5.001a2 2 0 1 1 4 0v5c0 .092-.006.183-.018.272A5.47 5.47 0 0 1 12 9.601V5a3 3 0 1 0-6 0v5a3 3 0 0 0 3 3Zm-4.5-3A4.5 4.5 0 0 0 9 14.5c0 .819.179 1.596.5 2.294v.706a.5.5 0 0 1-1 0v-2.022A5.5 5.5 0 0 1 3.5 10a.5.5 0 0 1 1 0ZM17 14.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm2 0a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-8 0a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Z"/>`),
		g.Group(children),
	)
}