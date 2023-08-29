package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 7.013V6a3.75 3.75 0 1 1 7.5 0v1.013c1.297.037 2.087.17 2.692.667c.83.68 1.06 1.834 1.523 4.143l.6 3c.664 3.32.996 4.98.096 6.079C19.76 22 18.067 22 14.68 22H9.32c-3.386 0-5.08 0-5.98-1.098c-.9-1.098-.568-2.758.096-6.079l.6-3c.462-2.309.693-3.463 1.522-4.143c.606-.496 1.396-.63 2.693-.667ZM9.75 6a2.25 2.25 0 0 1 4.5 0v1h-4.5V6ZM15 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}