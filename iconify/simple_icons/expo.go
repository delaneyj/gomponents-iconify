package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 20.084c.043.53.23 1.063.718 1.778c.58.849 1.576 1.315 2.303.567c.49-.505 5.794-9.776 8.35-13.29a.761.761 0 0 1 1.248 0c2.556 3.514 7.86 12.785 8.35 13.29c.727.748 1.723.282 2.303-.567c.57-.835.728-1.42.728-2.046c0-.426-8.26-15.798-9.092-17.078c-.8-1.23-1.044-1.498-2.397-1.542h-1.032c-1.353.044-1.597.311-2.398 1.542C8.267 3.991.33 18.758 0 19.77Z"/>`),
		g.Group(children),
	)
}