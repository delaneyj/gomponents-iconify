package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m12.684 3.27l.084.09l4.7 5.64H21a1 1 0 0 1 .991 1.131l-.02.112l-1.984 7.918c-.258 1.578-1.41 2.781-2.817 2.838L17 21l-10.148-.002c-1.37-.053-2.484-1.157-2.787-2.57l-.035-.185l-2-8a1 1 0 0 1 .857-1.237L3 9h3.53l4.702-5.64a1 1 0 0 1 1.452-.09zM12 12a3 3 0 0 0-2.98 2.65l-.015.174L9 15l.005.176A3 3 0 1 0 12 12zm0-6.438L9.135 9h5.73L12 5.562z"/></g>`),
		g.Group(children),
	)
}