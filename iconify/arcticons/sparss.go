package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.77 6.79a6.73 6.73 0 1 1 0 13.45h-6.72v-6.73a6.72 6.72 0 0 1 6.72-6.72Zm-19.67 0h.33a6.6 6.6 0 0 1 3.2.9A6.72 6.72 0 0 1 23 13.51h-4a6.73 6.73 0 1 1-13.45 0h4a6.72 6.72 0 0 1 3.35-5.82a6.6 6.6 0 0 1 3.2-.9Zm17.77 18.28a8.08 8.08 0 1 1-8.07 8.07a8.07 8.07 0 0 1 8.07-8.07ZM14.25 27l4.06 7l4.06 7H6.12l4.07-7l4.06-7Z"/>`),
		g.Group(children),
	)
}