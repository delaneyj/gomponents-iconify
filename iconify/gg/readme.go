package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Readme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 5.5h5a2 2 0 0 1 2 2v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Zm10 14c-.35 0-.687-.06-1-.17v.17a1 1 0 1 1-2 0v-.17c-.313.11-.65.17-1 .17H4a3 3 0 0 1-3-3v-10a3 3 0 0 1 3-3h5a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 15 3.5h5a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3h-6Zm-1-12v9a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-10a1 1 0 0 0-1-1h-5a2 2 0 0 0-2 2Zm-8 0h4v2H5v-2Zm10 0h4v2h-4v-2Zm4 3h-4v2h4v-2Zm-14 0h4v2H5v-2Zm14 3h-4v2h4v-2Zm-14 0h4v2H5v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}