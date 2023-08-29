package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NpmSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 2.885C1 1.844 1.844 1 2.885 1h10.23C14.156 1 15 1.844 15 2.885v10.23A1.885 1.885 0 0 1 13.115 15H2.885A1.885 1.885 0 0 1 1 13.115V2.885Zm1.885-.27a.27.27 0 0 0-.27.27v10.23c0 .15.12.27.27.27h5.653v-7a.808.808 0 1 1 1.616 0v7h2.961a.27.27 0 0 0 .27-.27V2.885a.27.27 0 0 0-.27-.27H2.885Z"/>`),
		g.Group(children),
	)
}