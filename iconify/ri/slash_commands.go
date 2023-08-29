package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashCommands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h2v4H3V3Zm6.788 18H7.66l6.551-18h2.129L9.788 21ZM21 3h-2v4h2V3Z"/>`),
		g.Group(children),
	)
}