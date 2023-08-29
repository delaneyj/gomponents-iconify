package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="20" cy="12" r="2" fill="currentColor"/><circle cx="4" cy="12" r="2" fill="currentColor"/><circle cx="12" cy="20" r="2" fill="currentColor"/><path fill="currentColor" d="m13.943 8.619l4.404-4.392l1.413 1.416l-4.405 4.392zM8.32 9.68l.31.32l1.42-1.41l-4.02-4.04h-.01l-.31-.32l-1.42 1.41l4.02 4.05zm7.09 4.26L14 15.35l3.99 4.01l.35.35l1.42-1.41l-3.99-4.01zm-6.82.01l-4.03 4.01l-.32.33l1.41 1.41l4.03-4.02l.33-.32z"/><circle cx="12" cy="4" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}