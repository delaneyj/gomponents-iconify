package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingChange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31 25a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 25 21a4.005 4.005 0 0 1 4 4zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 25 29a4.005 4.005 0 0 1-4-4zm-1.41-9L17 18.59L18.42 20l4-4l-4-4L17 13.41L19.59 16zm-9.18 0L13 13.41L11.58 12l-4 4l4 4L13 18.59L10.41 16z"/><path fill="currentColor" d="M4 9h22v7h2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h12v-2H4Zm0-5h22v3H4Z"/>`),
		g.Group(children),
	)
}