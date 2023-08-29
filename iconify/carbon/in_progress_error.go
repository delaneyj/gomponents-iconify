package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InProgressError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 24a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6Zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 24 20a4.005 4.005 0 0 1 4 4Zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 24 28a4.005 4.005 0 0 1-4-4Z"/><path fill="currentColor" d="M14 2a12 12 0 1 0 2 23.82v-2.022A10 10 0 1 1 14 4v10l4.343 4.343A7.975 7.975 0 0 1 24 16h1.82A11.93 11.93 0 0 0 14 2Z"/>`),
		g.Group(children),
	)
}