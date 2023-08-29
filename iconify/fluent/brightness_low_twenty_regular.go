package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessLowTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 3.5a.5.5 0 0 0-1 0v1a.5.5 0 0 0 1 0v-1ZM10 6a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-3 4a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm10 0a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1 0-1h1a.5.5 0 0 1 .5.5Zm-6.5 5.5a.5.5 0 0 0-1 0v1a.5.5 0 0 0 1 0v-1ZM5 10a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1 0-1h1a.5.5 0 0 1 .5.5Zm.854-4.854a.5.5 0 1 0-.708.708l.5.5a.5.5 0 1 0 .708-.708l-.5-.5Zm-.708 9.708a.5.5 0 0 0 .708 0l.5-.5a.5.5 0 0 0-.708-.708l-.5.5a.5.5 0 0 0 0 .708Zm9-9.708a.5.5 0 0 1 .708.708l-.5.5a.5.5 0 0 1-.708-.708l.5-.5Zm.708 9.708a.5.5 0 0 1-.708 0l-.5-.5a.5.5 0 0 1 .708-.708l.5.5a.5.5 0 0 1 0 .708Z"/>`),
		g.Group(children),
	)
}