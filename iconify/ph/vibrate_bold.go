package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrateBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M164 28H92a28 28 0 0 0-28 28v144a28 28 0 0 0 28 28h72a28 28 0 0 0 28-28V56a28 28 0 0 0-28-28Zm4 172a4 4 0 0 1-4 4H92a4 4 0 0 1-4-4V56a4 4 0 0 1 4-4h72a4 4 0 0 1 4 4Zm64-100v56a12 12 0 0 1-24 0v-56a12 12 0 0 1 24 0Zm-184 0v56a12 12 0 0 1-24 0v-56a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}