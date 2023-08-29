package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafariSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m4.958 10.042l1.906-3.178l3.178-1.906l-1.906 3.178l-3.178 1.906Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5A7.5 7.5 0 0 1 7.5 0A7.5 7.5 0 0 1 15 7.5A7.5 7.5 0 0 1 7.5 15A7.5 7.5 0 0 1 0 7.5Zm11.929-3.743a.5.5 0 0 0-.686-.686L6.136 6.136L3.07 11.243a.5.5 0 0 0 .686.686l5.107-3.065l3.065-5.107Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}