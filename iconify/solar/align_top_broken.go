package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M10 3h12M2 3h4M5 13v5.5c0 .935 0 1.402.201 1.75a1.5 1.5 0 0 0 .549.549C6.098 21 6.565 21 7.5 21s1.402 0 1.75-.201a1.5 1.5 0 0 0 .549-.549c.201-.348.201-.815.201-1.75v-9c0-.935 0-1.402-.201-1.75a1.5 1.5 0 0 0-.549-.549C8.902 7 8.435 7 7.5 7s-1.402 0-1.75.201a1.5 1.5 0 0 0-.549.549c-.161.28-.193.636-.2 1.25"/><path d="M16.5 7c-.935 0-1.402 0-1.75.201a1.5 1.5 0 0 0-.549.549C14 8.098 14 8.565 14 9.5v6c0 .935 0 1.402.201 1.75a1.5 1.5 0 0 0 .549.549c.348.201.815.201 1.75.201s1.402 0 1.75-.201a1.5 1.5 0 0 0 .549-.549c.201-.348.201-.815.201-1.75v-6c0-.935 0-1.402-.201-1.75a1.5 1.5 0 0 0-.549-.549C17.902 7 17.435 7 16.5 7Z"/></g>`),
		g.Group(children),
	)
}