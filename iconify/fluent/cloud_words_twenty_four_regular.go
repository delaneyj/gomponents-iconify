package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWordsTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.505 10.785a4.5 4.5 0 0 1 8.99 0a.75.75 0 0 0 .75.715h.255a3 3 0 1 1 0 6h-11a3 3 0 1 1 0-6h.256c.4 0 .73-.315.749-.715ZM12 5a6.001 6.001 0 0 0-5.92 5.02A4.5 4.5 0 0 0 6.5 19h11a4.5 4.5 0 0 0 .42-8.98A6.001 6.001 0 0 0 12 5Zm-2 4.5a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Zm-4.5 4.75a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Zm8.25-.75a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Z"/>`),
		g.Group(children),
	)
}