package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12.707 4.293a1 1 0 0 0-1.414 0l-4 4A1 1 0 0 0 8 10h8a1 1 0 0 0 .707-1.707l-4-4zM8 14a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 14H8z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}