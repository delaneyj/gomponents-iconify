package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.403 12.652a3 3 0 0 0 0-5.304a3 3 0 0 0-3.75-3.751a3 3 0 0 0-5.305 0a3 3 0 0 0-3.751 3.75a3 3 0 0 0 0 5.305a3 3 0 0 0 3.75 3.751a3 3 0 0 0 5.305 0a3 3 0 0 0 3.751-3.75Zm-2.546-4.46a.75.75 0 0 0-1.214-.883l-3.483 4.79l-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089l4-5.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}