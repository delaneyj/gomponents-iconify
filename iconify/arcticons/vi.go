package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="9.128" height="24.342" x="33.872" y="8.251" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".507"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.018 8.635L18.55 26.507L14.082 8.635a.507.507 0 0 0-.492-.384H5.507a.507.507 0 0 0-.492.63l5.832 23.327a.507.507 0 0 0 .492.384h14.422a.507.507 0 0 0 .492-.384l5.832-23.327a.507.507 0 0 0-.492-.63H23.51a.507.507 0 0 0-.492.384Z"/><circle cx="38.436" cy="39.185" r="4.564" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}