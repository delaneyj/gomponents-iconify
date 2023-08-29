package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shelter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 20v-2.5s-1.5-1-4-1s-4 1-4 1V20m15 0v-2.5s-1.5-1-4-1a8.97 8.97 0 0 0-1.5.124m8.5-6.374l-.247-.113a20 20 0 0 1-8.942-8.104L13 1.5h-2l-.311.533a20 20 0 0 1-8.942 8.104l-.247.113V22.5h21V10.25ZM8.35 14.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Zm7 0s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}