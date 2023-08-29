package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.05 2.05a2.801 2.801 0 0 0 3.136.571A2.801 2.801 0 0 0 17 0a2.8 2.8 0 0 0 1.814 2.621a2.801 2.801 0 0 0 3.136-.57a2.8 2.8 0 0 0-.571 3.135A2.801 2.801 0 0 0 24 7M3 11.5h2m16.5-3h-3a3 3 0 0 1-3-3h-6a3 3 0 0 1-3 3h-6v13h18a3 3 0 0 0 3-3v-10Zm-5 6a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}