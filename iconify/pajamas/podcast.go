package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Podcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.397 2.59a6.5 6.5 0 0 1 8.227 9.978a.75.75 0 0 0 1.067 1.054a8 8 0 1 0-11.382 0a.75.75 0 1 0 1.067-1.054A6.5 6.5 0 0 1 4.397 2.59ZM8 4.5a3.5 3.5 0 0 1 2.5 5.949a.75.75 0 1 0 1.072 1.05a5 5 0 1 0-7.144 0A.75.75 0 0 0 5.5 10.45A3.5 3.5 0 0 1 8 4.5ZM10 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-1.25 4.25a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}