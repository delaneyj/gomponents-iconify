package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaImage2Fill0"><g id="evaImage2Fill1"><path id="evaImage2Fill2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM8 7a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 8 7Zm11 10.83A1.09 1.09 0 0 1 18 19H6l7.57-6.82a.69.69 0 0 1 .93 0l4.5 4.48Z"/></g></g>`),
		g.Group(children),
	)
}