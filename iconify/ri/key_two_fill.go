package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.313 11.566l7.94-7.94l2.121 2.12l-1.414 1.415l2.121 2.121l-3.535 3.536l-2.121-2.121l-2.99 2.99a5.002 5.002 0 0 1-7.97 5.848a5 5 0 0 1 5.848-7.97Zm-.899 5.848a2 2 0 1 0-2.828-2.828a2 2 0 0 0 2.828 2.828Z"/>`),
		g.Group(children),
	)
}