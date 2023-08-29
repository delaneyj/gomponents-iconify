package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.484 0l-7.516 8.719L.702 6.292L12.28 16.12L.895 26.073l11.214-2.646L19.791 32l.948-11.469l10.557-4.646l-10.62-4.438L19.494-.001z"/>`),
		g.Group(children),
	)
}