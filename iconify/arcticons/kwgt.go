package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kwgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="8.21" height="39" x="10.173" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.642"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.006 4.5h-.821a6.568 6.568 0 0 0-6.569 6.568v4.358a.82.82 0 0 1-.269.608l-7.426 6.75a1.642 1.642 0 0 0 0 2.43l7.426 6.752a.82.82 0 0 1 .27.608v4.357a6.568 6.568 0 0 0 6.568 6.569h.82a.821.821 0 0 0 .822-.821V29.606a1.642 1.642 0 0 0-.481-1.16L33.48 24.58a.821.821 0 0 1 0-1.162l3.865-3.864a1.642 1.642 0 0 0 .48-1.161V5.32a.821.821 0 0 0-.82-.821Z"/>`),
		g.Group(children),
	)
}