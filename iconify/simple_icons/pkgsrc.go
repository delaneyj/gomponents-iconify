package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pkgsrc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.908 8.763l9.157-5.132L11.25 0L1.62 4.42Zm1.5 2.29l9-5.368l-.948 11.84l-8.191 6.382zM.593 6.712L1.619 18.79L11.922 24l-.12-12.788Z"/>`),
		g.Group(children),
	)
}