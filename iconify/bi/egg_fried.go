package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggFried(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/><path d="M13.997 5.17a5 5 0 0 0-8.101-4.09A5 5 0 0 0 1.28 9.342a5 5 0 0 0 8.336 5.109a3.5 3.5 0 0 0 5.201-4.065a3.001 3.001 0 0 0-.822-5.216zm-1-.034a1 1 0 0 0 .668.977a2.001 2.001 0 0 1 .547 3.478a1 1 0 0 0-.341 1.113a2.5 2.5 0 0 1-3.715 2.905a1 1 0 0 0-1.262.152a4 4 0 0 1-6.67-4.087a1 1 0 0 0-.2-1a4 4 0 0 1 3.693-6.61a1 1 0 0 0 .8-.2a4 4 0 0 1 6.48 3.273z"/></g>`),
		g.Group(children),
	)
}