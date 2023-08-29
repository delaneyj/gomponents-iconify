package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.025 8c0 2.485-2.02 4.5-4.513 4.5A4.506 4.506 0 0 1 0 8c0-2.486 2.02-4.5 4.512-4.5A4.506 4.506 0 0 1 9.025 8zm4.95 0c0 2.34-1.01 4.236-2.256 4.236c-1.246 0-2.256-1.897-2.256-4.236c0-2.34 1.01-4.236 2.256-4.236c1.246 0 2.256 1.897 2.256 4.236zM16 8c0 2.096-.355 3.795-.794 3.795c-.438 0-.793-1.7-.793-3.795c0-2.096.355-3.795.794-3.795c.438 0 .793 1.699.793 3.795z"/>`),
		g.Group(children),
	)
}