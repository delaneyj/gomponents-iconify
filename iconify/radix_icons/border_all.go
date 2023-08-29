package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M.25 1A.75.75 0 0 1 1 .25h13a.75.75 0 0 1 .75.75v13a.75.75 0 0 1-.75.75H1A.75.75 0 0 1 .25 14V1Zm1.5.75v11.5h11.5V1.75H1.75Z" clip-rule="evenodd"/><rect width="1" height="1" x="7" y="5" rx=".5"/><rect width="1" height="1" x="7" y="3" rx=".5"/><rect width="1" height="1" x="7" y="7" rx=".5"/><rect width="1" height="1" x="5" y="7" rx=".5"/><rect width="1" height="1" x="3" y="7" rx=".5"/><rect width="1" height="1" x="9" y="7" rx=".5"/><rect width="1" height="1" x="11" y="7" rx=".5"/><rect width="1" height="1" x="7" y="9" rx=".5"/><rect width="1" height="1" x="7" y="11" rx=".5"/></g>`),
		g.Group(children),
	)
}