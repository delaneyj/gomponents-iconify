package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agoradesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.714 9.284c5.168-3.136 17.493-6.935 29.729-1.026c-8.57 2.087-16.256 16.31-4.285 25.808C18.107 39.102 5.561 23.553 9.714 9.284Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.66 38.175c-.873 1.59-1.328 2.425-3.128 3.467c-2.075 1.2-6.756 1.178-8.857 1.196c-6.166.054-18.606-.379-15.679-12.284c4.362 7.886 16.333 12.502 27.664 7.62Z"/>`),
		g.Group(children),
	)
}