package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramerLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.382 1.296A.5.5 0 0 1 3.84.997h7.66a.5.5 0 0 1 .5.5V5.5a.5.5 0 0 1-.5.5H8.635l2.894 3.162a.5.5 0 0 1-.369.838H8v3.5a.5.5 0 0 1-.854.354l-4-4A.5.5 0 0 1 3 9.5v-4a.5.5 0 0 1 .5-.5h2.865L3.471 1.835a.5.5 0 0 1-.089-.54ZM7.72 5L4.975 1.997H11V5H7.72Zm-.44 1H4v3h6.025L7.28 6Zm-2.573 4L7 12.293V10H4.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}