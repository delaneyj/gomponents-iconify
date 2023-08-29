package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M14 20.95h6V10.78L8 7.34V3.05H4v17.9h6v-5.64h4v5.64z" fill="currentColor"/>`),
		g.Group(children),
	)
}