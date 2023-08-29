package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectObjectTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1-.25c0 .414.336.75.75.75h6.5a.75.75 0 0 0 0-1.5h-6.5a.75.75 0 0 0-.75.75ZM4.75 8a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0v-6.5A.75.75 0 0 0 4.75 8Zm14.5 0a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0v-6.5a.75.75 0 0 0-.75-.75ZM8.75 20a.75.75 0 0 1 0-1.5h6.5a.75.75 0 0 1 0 1.5h-6.5ZM5 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM21 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}