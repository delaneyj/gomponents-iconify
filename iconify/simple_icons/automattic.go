package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automattic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.521 8.11a1.497 1.497 0 0 1 .433 2.102l-3.511 5.441a1.496 1.496 0 0 1-2.068.457a1.507 1.507 0 0 1-.44-2.08l3.513-5.44a1.5 1.5 0 0 1 .943-.655c.39-.085.796-.04 1.13.175zM11.98 23.03C4.713 23.03 0 17.79 0 12.338v-.676C0 6.117 4.713.97 11.98.97C19.246.97 24 6.117 24 11.662v.676c0 5.453-4.713 10.692-12.02 10.692zm8.133-11.31c0-3.974-2.888-7.51-8.133-7.51c-5.245 0-8.087 3.542-8.087 7.51v.497c0 3.974 2.888 7.578 8.087 7.578c5.198 0 8.133-3.604 8.133-7.578v-.497z"/>`),
		g.Group(children),
	)
}