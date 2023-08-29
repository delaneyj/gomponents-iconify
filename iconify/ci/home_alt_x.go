package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-9.643a1.01 1.01 0 0 1 .293-.65l7-7a1 1 0 0 1 1.415 0l7 7a.994.994 0 0 1 .292.707V21a1 1 0 0 1-1 1ZM12 5.828l-6 6V20h12v-8.172l-6-6ZM9.878 18.071l-1.413-1.414l2.121-2.121l-2.121-2.122L9.879 11L12 13.121L14.121 11l1.414 1.414l-2.121 2.121l2.121 2.121l-1.413 1.413L12 15.949l-2.121 2.122h-.001Z"/>`),
		g.Group(children),
	)
}