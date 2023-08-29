package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OctagonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M15.3 2H8.7c-.562 0-1.016.201-1.407.593l-4.7 4.7A1.894 1.894 0 0 0 2 8.7v6.6c0 .562.201 1.016.593 1.407l4.7 4.7c.391.392.845.593 1.407.593h6.6c.562 0 1.016-.201 1.407-.593l4.7-4.7c.392-.391.593-.845.593-1.407V8.7c0-.562-.201-1.016-.593-1.407l-4.7-4.7A1.894 1.894 0 0 0 15.3 2z"/></g>`),
		g.Group(children),
	)
}