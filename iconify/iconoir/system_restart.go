package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemRestart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2v4m0 12v4m10-10h-4M6 12H2m2.929-7.071l2.828 2.828m8.486 8.486l2.828 2.828m0-14.142l-2.828 2.828m-8.486 8.486L4.93 19.07"/>`),
		g.Group(children),
	)
}