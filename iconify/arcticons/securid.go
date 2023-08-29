package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Securid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="35.817" cy="27.952" r="7.683" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.44 20.45a8.723 8.723 0 0 0-16.68-2.882a6.661 6.661 0 0 0-2.998-.734a6.733 6.733 0 0 0-6.586 5.341a6.73 6.73 0 0 0 .058 13.461h24.583"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.109 23.567l-6.929 6.928l-2.933-2.934"/>`),
		g.Group(children),
	)
}