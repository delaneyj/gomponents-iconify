package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ankh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 13h-4.807A7.783 7.783 0 0 0 17 6.9a5.001 5.001 0 0 0-10 0a7.783 7.783 0 0 0 3.307 6.1H5.5a.5.5 0 0 0 0 1h6v7.5a.5.5 0 0 0 1 0V14h6a.5.5 0 0 0 0-1zm-6.501-.07C11.132 12.436 8 10.399 8 6.9a4.001 4.001 0 0 1 8 0c0 3.488-3.134 5.533-4.001 6.03z"/>`),
		g.Group(children),
	)
}