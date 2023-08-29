package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.998 1a1 1 0 0 1 1 1v20a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14Zm-1 11h-12v9h12v-9Zm-8 2v4h-2v-4h2Zm8-11h-12v7h12V3Zm-8 2v3h-2V5h2Z"/>`),
		g.Group(children),
	)
}