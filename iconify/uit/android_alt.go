package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 9a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5zm12.344-4.602l1.073-1.622a.5.5 0 0 0-.834-.552l-1.06 1.602a5.946 5.946 0 0 0-6.046 0l-1.06-1.602a.5.5 0 0 0-.834.552l1.073 1.622A5.988 5.988 0 0 0 6 9v8.5a.5.5 0 0 0 .5.5h3v3.5a.5.5 0 0 0 1 0V18h3v3.5a.5.5 0 0 0 1 0V18h3a.5.5 0 0 0 .5-.5V9a5.99 5.99 0 0 0-2.156-4.602zM17 17H7v-7h10v7zM7 9a5 5 0 0 1 10 0H7zm13.5 0a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 1 0v-5a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}