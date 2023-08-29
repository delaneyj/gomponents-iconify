package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 12a2 2 0 0 1 2-2h16a2 2 0 0 1 1.515 3.305c-1.76 2.044-4.362 5.723-6.678 9.977c-2.327 4.272-4.267 8.946-4.858 13.006a2 2 0 0 1-3.958-.576c.69-4.74 2.882-9.899 5.302-14.344a73.401 73.401 0 0 1 4.6-7.368H16a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}