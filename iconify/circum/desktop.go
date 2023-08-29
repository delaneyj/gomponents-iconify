package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.435 3.065H5.565a2.5 2.5 0 0 0-2.5 2.5v8.87a2.5 2.5 0 0 0 2.5 2.5h2.91l-.37 3H7a.5.5 0 0 0 0 1h10.01a.5.5 0 0 0 0-1H15.9l-.37-3h2.91a2.5 2.5 0 0 0 2.5-2.5v-8.87a2.5 2.5 0 0 0-2.505-2.5Zm-9.33 16.87l.38-3h5.03l.37 3Zm10.83-5.5a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-.5h15.87Zm0-1.5H4.065v-7.37a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}