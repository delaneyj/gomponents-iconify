package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 25a6.99 6.99 0 0 1-5.833-3.129l1.666-1.107a5 5 0 0 0 8.334 0l1.666 1.107A6.99 6.99 0 0 1 20 25zm4-11a2 2 0 1 0 2 2a1.98 1.98 0 0 0-2-2zm-8 0a2 2 0 1 0 2 2a1.98 1.98 0 0 0-2-2z"/><path fill="currentColor" d="M28 8h-6V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v10a10.01 10.01 0 0 0 8.803 9.921A9.999 9.999 0 0 0 30 20V10a2.002 2.002 0 0 0-2-2ZM4 14V4h16v4h-8a2.002 2.002 0 0 0-2 2v10a9.963 9.963 0 0 0 .168 1.78A8.008 8.008 0 0 1 4 14Zm24 6a8 8 0 0 1-16 0V10h16Z"/>`),
		g.Group(children),
	)
}