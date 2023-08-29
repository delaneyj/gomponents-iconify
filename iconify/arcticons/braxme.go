package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braxme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.616 40.508A19.5 19.5 0 1 1 38.92 36.56m0 0c.92 2.404 1.026 5.284-.234 6.705c-1.856 1.542-6.499-5.093-11.107-3.653c-5.238 2.87-11.135 2.76-13.962.896"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.172 23.967a4.91 4.91 0 0 1 0 9.821h-8.103V14.146h8.103a4.91 4.91 0 0 1 0 9.821Zm0 0h-8.103"/>`),
		g.Group(children),
	)
}