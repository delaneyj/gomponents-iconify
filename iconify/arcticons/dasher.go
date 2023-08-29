package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dasher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.224 31.315H28.5a4.113 4.113 0 0 0 4.114-4.114h0a4.113 4.113 0 0 0-4.114-4.113H15.386"/><rect width="29.99" height="32.597" x="9.005" y="10.903" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.384"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.542 10.903V7.268A2.768 2.768 0 0 0 27.774 4.5h-7.548a2.768 2.768 0 0 0-2.768 2.768v3.635"/>`),
		g.Group(children),
	)
}