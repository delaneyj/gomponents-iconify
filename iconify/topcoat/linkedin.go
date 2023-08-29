package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M6.5 17.5h6v17h-6v-17zm20.141-.58c1.689.01 3.85.84 5.17 2.18c1.369 1.53 1.689 3.259 1.689 5.4v10h-6v-9.29c0-1.7-.689-3.771-2.939-3.84c-1.32.021-2.17.78-2.801 2.06c-.18.42-.14.891-.14 1.36l-.12 9.71h-6v-17h6l.12 2.22a6.19 6.19 0 0 1 1.69-1.829c.96-.691 2.081-.952 3.331-.971zM9.5 9.35c1.54.021 3.07 1.23 3.13 3.07c.04 1.641-1.39 3.07-3.17 3.07h-.04c-1.53 0-3.03-1.25-3.1-3.07c.03-1.62 1.39-3.029 3.18-3.07zM.5 1.5v39h39v-39H.5z"/>`),
		g.Group(children),
	)
}