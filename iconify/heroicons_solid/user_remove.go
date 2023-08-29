package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm3 11a6 6 0 0 0-12 0h12Zm-1-9a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4Z"/>`),
		g.Group(children),
	)
}