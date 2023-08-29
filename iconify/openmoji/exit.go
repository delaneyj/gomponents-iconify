package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiExit0" d="M40.854 21.275h10.85v10.85h-10.85z"/><path id="openmojiExit1" d="M35 37.979V18.104H15.125v39.75h39.75V38.012l-4.562-.007l-.282-.001z"/></defs><g fill="#9B9B9A"><use href="#openmojiExit0"/><use href="#openmojiExit1"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiExit1"/><use href="#openmojiExit0"/><path d="m46.279 26.7l13.153-13.153l.558-.557V21M52 13h8"/></g>`),
		g.Group(children),
	)
}