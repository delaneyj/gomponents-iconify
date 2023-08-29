package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M68 60a68 68 0 1 0 68 68a68.07 68.07 0 0 0-68-68Zm0 112a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44ZM184 60c-23.63 0-36 34.21-36 68s12.37 68 36 68s36-34.21 36-68s-12.37-68-36-68Zm0 111.87c-3.74-2.16-12-17.09-12-43.87s8.26-41.71 12-43.87c3.74 2.16 12 17.09 12 43.87s-8.26 41.71-12 43.87ZM256 72v112a12 12 0 0 1-24 0V72a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}