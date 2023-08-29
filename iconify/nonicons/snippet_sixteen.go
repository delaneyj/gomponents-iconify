package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnippetSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 4c0-.966.784-1.75 1.75-1.75h10c.966 0 1.75.784 1.75 1.75v6.5a.75.75 0 0 1-1.5 0V4a.25.25 0 0 0-.25-.25h-10A.25.25 0 0 0 3 4v6.5a.75.75 0 0 1-1.5 0V4Z" clip-rule="evenodd"/><path fill="currentColor" d="M2.25 13.5a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm12 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-3 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-6 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm3 0a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}