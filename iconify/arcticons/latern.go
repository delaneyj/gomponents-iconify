package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Latern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.021 14.417c.426 7.968 4.38 28.945 9.72 29.082c5.353.138 7.55-15.6 10.237-29.157l-5.341-1.929l.074-2.893l-4.896-2.597l-4.897 2.968l.074 2.893Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.332 17.863a25.428 25.428 0 0 1 19.026-.324M23.888 43.5l.012-27.646m-2.104-7.862a2.328 2.328 0 1 1 4.032 0"/>`),
		g.Group(children),
	)
}