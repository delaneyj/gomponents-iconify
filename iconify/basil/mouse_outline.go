package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.053 2.407a6.75 6.75 0 0 1 5.697 6.667v5.852a6.75 6.75 0 0 1-13.5 0V9.074a6.75 6.75 0 0 1 7.803-6.667Zm4.197 6.667a5.25 5.25 0 0 0-4.5-5.196V9.25h4.5v-.176Zm-6 .176h-4.5v-.176a5.25 5.25 0 0 1 4.5-5.196V9.25Zm-4.5 5.676V10.75h10.5v4.176a5.25 5.25 0 0 1-10.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}