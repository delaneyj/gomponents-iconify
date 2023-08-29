package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18.75H4a.76.76 0 0 1-.65-.37a.77.77 0 0 1 0-.75l8-14a.78.78 0 0 1 1.3 0l8 14a.77.77 0 0 1 0 .75a.76.76 0 0 1-.65.37Zm-14.71-1.5h13.42L12 5.51Z"/><path fill="currentColor" d="M12 13.25a.76.76 0 0 1-.75-.75V9a.75.75 0 0 1 1.5 0v3.5a.76.76 0 0 1-.75.75Zm0 3a.76.76 0 0 1-.75-.75V15a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}