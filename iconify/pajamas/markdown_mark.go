package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.308 5.308v5.23h1.538v-3l1.539 1.924l1.538-1.924v3h1.539v-5.23H6.923L5.385 7.23L3.846 5.308H2.308ZM9.615 8l2.308 2.539L14.231 8h-1.539V5.308h-1.538V8H9.615Z"/><path fill="currentColor" fill-rule="evenodd" d="M1.154 3C.517 3 0 3.517 0 4.154v7.538c0 .637.517 1.154 1.154 1.154h13.692c.637 0 1.154-.517 1.154-1.154V4.154C16 3.517 15.483 3 14.846 3H1.154ZM.769 4.154c0-.213.172-.385.385-.385h13.692c.213 0 .385.172.385.385v7.538a.385.385 0 0 1-.385.385H1.154a.385.385 0 0 1-.385-.385V4.154Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}