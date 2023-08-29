package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HybridNetworkingAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 22a3.958 3.958 0 0 0-2.02.566L17.414 16l6.567-6.567A3.952 3.952 0 0 0 26 10a4 4 0 1 0-4-4a3.951 3.951 0 0 0 .567 2.019L16 14.586L9.434 8.02A3.958 3.958 0 0 0 10 6a4 4 0 1 0-4 4a3.958 3.958 0 0 0 2.02-.566L14.586 16l-6.567 6.567A3.952 3.952 0 0 0 6 22a4 4 0 1 0 4 4a3.951 3.951 0 0 0-.567-2.019L16 17.414l6.566 6.566A3.958 3.958 0 0 0 22 26a4 4 0 1 0 4-4Zm0-18a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2ZM6 28a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}