package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.9 7.5a6.6 6.6 0 1 1 13.2 0a6.6 6.6 0 0 1-13.2 0Zm6.6-5.7a5.7 5.7 0 1 0 0 11.4a5.7 5.7 0 0 0 0-11.4ZM3.075 7.5a4.425 4.425 0 1 1 8.85 0a4.425 4.425 0 0 1-8.85 0ZM7.5 3.925a3.575 3.575 0 1 0 0 7.15a3.575 3.575 0 0 0 0-7.15Zm0 1.325a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM6.05 7.5a1.45 1.45 0 1 1 2.9 0a1.45 1.45 0 0 1-2.9 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}