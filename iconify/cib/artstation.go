package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m0 23.63l2.703 4.672a3.232 3.232 0 0 0 2.885 1.781h17.943l-3.724-6.453zm32 .031a3.25 3.25 0 0 0-.516-1.75L20.968 3.635a3.225 3.225 0 0 0-2.854-1.719h-5.557l16.24 28.135l2.563-4.432c.5-.849.641-1.224.641-1.958zm-14.839-4.614L9.906 6.479l-7.26 12.568z"/>`),
		g.Group(children),
	)
}