package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBorderThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.232 16.232L10 20.465V10h12v10.465l-4.232-4.233a2.5 2.5 0 0 0-3.536 0ZM19 11a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-2.646 6.646L20.707 22h-9.414l4.353-4.354a.5.5 0 0 1 .708 0ZM3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5v17a4.5 4.5 0 0 1-4.5 4.5h-17A4.5 4.5 0 0 1 3 24.5v-17ZM9 8a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1H9Z"/>`),
		g.Group(children),
	)
}