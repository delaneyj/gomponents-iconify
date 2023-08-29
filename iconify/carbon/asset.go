package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 24a4 4 0 1 1 4-4a4.004 4.004 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M30 6a4.004 4.004 0 0 0-4-4a3.949 3.949 0 0 0-1.854.477l-16.389 8.48a9.992 9.992 0 1 0 13.309 13.236l8.49-16.4A3.95 3.95 0 0 0 30 6Zm-4-2a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2Zm-3.98 1.808c-.004.065-.02.126-.02.192a4.004 4.004 0 0 0 4 4c.065 0 .124-.016.188-.019l-4.332 8.362a10.017 10.017 0 0 0-8.215-8.196ZM12 28a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}