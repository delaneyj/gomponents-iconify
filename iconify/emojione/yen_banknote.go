package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#89664c" d="M0 46.5h64v9H0z"/><path fill="#d3976e" d="M0 8.5h64v38H0z"/><path fill="#94989b" d="M24 46.5h16v9H24z"/><path fill="#89664c" d="M4 12.5h56v30H4z"/><path fill="#d3976e" d="M7 15.5h50v24H7z"/><circle cx="45" cy="27.5" r="8" fill="#89664c"/><path fill="#d0d0d0" d="M24 8.5h16v38H24z"/><path fill="#fff" d="m22 21.2l-1.2-1.7l-4.8 4.6l-4.8-4.6l-1.2 1.7l4.5 4.3h-2.3v2.2H15V30h-2.8v2.2H15v3.3h2v-3.3h2.8V30H17v-2.3h2.8v-2.2h-2.3z"/>`),
		g.Group(children),
	)
}