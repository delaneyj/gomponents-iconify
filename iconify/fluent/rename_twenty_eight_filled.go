package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RenameTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.75 2a.75.75 0 0 0 0 1.5h1.5v21h-1.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-1.5v-21h1.5a.75.75 0 0 0 0-1.5h-4.5Zm-5.5 4.02h6v16h-6A3.25 3.25 0 0 1 3 18.77v-9.5a3.25 3.25 0 0 1 3.25-3.25Zm15.5 16h-6v-16h6A3.25 3.25 0 0 1 25 9.27v9.5a3.25 3.25 0 0 1-3.25 3.25Z"/>`),
		g.Group(children),
	)
}