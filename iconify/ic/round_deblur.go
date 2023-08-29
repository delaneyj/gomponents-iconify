package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundDeblur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3v18a9 9 0 0 0 0-18z"/><circle cx="6" cy="14" r="1" fill="currentColor"/><circle cx="6" cy="18" r="1" fill="currentColor"/><circle cx="6" cy="10" r="1" fill="currentColor"/><circle cx="3" cy="10" r=".5" fill="currentColor"/><circle cx="6" cy="6" r="1" fill="currentColor"/><circle cx="3" cy="14" r=".5" fill="currentColor"/><circle cx="10" cy="21" r=".5" fill="currentColor"/><circle cx="10" cy="3" r=".5" fill="currentColor"/><circle cx="10" cy="6" r="1" fill="currentColor"/><circle cx="10" cy="14" r="1.5" fill="currentColor"/><circle cx="10" cy="10" r="1.5" fill="currentColor"/><circle cx="10" cy="18" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}