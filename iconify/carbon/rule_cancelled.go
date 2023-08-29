package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleCancelled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 24a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 24 20a4.005 4.005 0 0 1 4 4zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 24 28a4.005 4.005 0 0 1-4-4zM8 16h10v2H8zm0-6h12v2H8z"/><path fill="currentColor" d="m14 27.733l-5.234-2.79A8.986 8.986 0 0 1 4 17V4h20v11h2V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v13a10.981 10.981 0 0 0 5.824 9.707L14 30Z"/>`),
		g.Group(children),
	)
}