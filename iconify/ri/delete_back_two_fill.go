package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteBackTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.535 3h14.464a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H6.535a1 1 0 0 1-.833-.445l-5.333-8a1 1 0 0 1 0-1.11l5.333-8A1 1 0 0 1 6.535 3Zm6.464 7.586l-2.828-2.829l-1.414 1.415L11.585 12l-2.828 2.828l1.414 1.415l2.828-2.829l2.829 2.829l1.414-1.415L14.413 12l2.829-2.828l-1.414-1.415l-2.829 2.829Z"/>`),
		g.Group(children),
	)
}