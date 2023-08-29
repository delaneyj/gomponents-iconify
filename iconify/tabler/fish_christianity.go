package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishChristianity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 7S16.354 17 9.692 17c-3.226.025-6.194-1.905-7.692-5c1.498-3.095 4.466-5.025 7.692-5C16.354 7 22 17 22 17"/>`),
		g.Group(children),
	)
}