package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortByAlphabetLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 7H3m7 5H3m5 5H3"/><path fill="currentColor" d="M11.316 16.692a.75.75 0 1 0 1.368.616l-1.368-.616ZM16.5 7l.684-.308a.75.75 0 0 0-1.368 0L16.5 7Zm3.816 10.308a.75.75 0 1 0 1.368-.616l-1.368.616Zm-.952-3.944l.684-.308l-.684.308Zm-5.728-.75a.75.75 0 0 0 0 1.5v-1.5Zm-.952 4.694l4.5-10l-1.368-.616l-4.5 10l1.368.616Zm9-.616l-1.636-3.636l-1.368.615l1.636 3.637l1.368-.616Zm-1.636-3.636l-2.864-6.364l-1.368.616l2.864 6.363l1.368-.615Zm-.684-.442h-5.728v1.5h5.728v-1.5Z"/></g>`),
		g.Group(children),
	)
}