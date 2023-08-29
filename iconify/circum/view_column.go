package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.44 3.06H5.56a2.507 2.507 0 0 0-2.5 2.5v12.88a2.514 2.514 0 0 0 2.5 2.5h12.88a2.514 2.514 0 0 0 2.5-2.5V5.56a2.507 2.507 0 0 0-2.5-2.5ZM8.67 19.94H5.56a1.511 1.511 0 0 1-1.5-1.5V5.56a1.5 1.5 0 0 1 1.5-1.5h3.11Zm1-15.88h4.66v15.88H9.67Zm10.27 14.38a1.511 1.511 0 0 1-1.5 1.5h-3.11V4.06h3.11a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}