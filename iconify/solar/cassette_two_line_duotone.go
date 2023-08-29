package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CassetteTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4h4c3.771 0 5.657 0 6.828 1.172C22 6.343 22 8.229 22 12c0 3.771 0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12Z"/><path fill="currentColor" d="M8.5 10.75a.75.75 0 0 0 0 1.5v-1.5Zm7.5 1.5a.75.75 0 0 0 0-1.5v1.5Zm-7.5 3a.75.75 0 0 0 0 1.5v-1.5Zm7.5 1.5a.75.75 0 0 0 0-1.5v1.5Zm1.25-3a1.5 1.5 0 0 1-1.5 1.5v1.5a3 3 0 0 0 3-3h-1.5Zm-1.5-1.5a1.5 1.5 0 0 1 1.5 1.5h1.5a3 3 0 0 0-3-3v1.5Zm-7.5 3a1.5 1.5 0 0 1-1.5-1.5h-1.5a3 3 0 0 0 3 3v-1.5Zm0-4.5a3 3 0 0 0-3 3h1.5a1.5 1.5 0 0 1 1.5-1.5v-1.5Zm.25 1.5H16v-1.5H8.5v1.5Zm0 4.5H16v-1.5H8.5v1.5Z"/><path fill="currentColor" d="M15.75 15.25a1.5 1.5 0 0 1-1.5-1.5h-1.5a3 3 0 0 0 3 3v-1.5Zm0-4.5a3 3 0 0 0-3 3h1.5a1.5 1.5 0 0 1 1.5-1.5v-1.5Zm-6 3a1.5 1.5 0 0 1-1.5 1.5v1.5a3 3 0 0 0 3-3h-1.5Zm-1.5-1.5a1.5 1.5 0 0 1 1.5 1.5h1.5a3 3 0 0 0-3-3v1.5Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.5 4.5l-.527 1.404c-.47 1.256-.706 1.884-1.22 2.24c-.514.356-1.184.356-2.525.356h-2.456c-1.34 0-2.011 0-2.525-.356c-.514-.356-.75-.984-1.22-2.24L6.5 4.5" opacity=".5"/></g>`),
		g.Group(children),
	)
}