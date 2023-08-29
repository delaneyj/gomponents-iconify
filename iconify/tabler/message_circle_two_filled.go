package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCircleTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M5.821 4.91c3.898-2.765 9.469-2.539 13.073.536c3.667 3.127 4.168 8.238 1.152 11.897c-2.842 3.447-7.965 4.583-12.231 2.805l-.232-.101l-4.375.931l-.075.013l-.11.009l-.113-.004l-.044-.005l-.11-.02l-.105-.034l-.1-.044l-.076-.042l-.108-.077l-.081-.074l-.073-.083l-.053-.075l-.065-.115l-.042-.106l-.031-.113l-.013-.075l-.009-.11l.004-.113l.005-.044l.02-.11l.022-.072l1.15-3.451l-.022-.036C.969 12.45 1.97 7.805 5.59 5.079l.23-.168z"/></g>`),
		g.Group(children),
	)
}