package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.149 19.17C19.834 22.121 16.268 24 12.263 24c-.097 0-.194-.001-.29-.003h.014a12.513 12.513 0 0 1-10.134-4.8l-.021-.027L-.001 21v-5.994h5.994l-1.644 1.646c1.331 2.188 3.52 3.742 6.089 4.197l.054.008v-9.061c-2.606-.689-4.495-3.026-4.495-5.803a5.994 5.994 0 1 1 7.535 5.794l-.042.009v9.062a8.965 8.965 0 0 0 6.121-4.166l.022-.039l-1.652-1.647h5.994V21zM14.988 6.016a2.997 2.997 0 1 0-5.995-.001a2.997 2.997 0 0 0 5.995.002z"/>`),
		g.Group(children),
	)
}