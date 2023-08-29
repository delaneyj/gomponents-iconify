package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qunar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 20.573s-2.656-1.255-3.59 0c-2.864 3.85-18.707 9.96-16.625-4.805a12.271 12.271 0 0 0-6.738-.055c-.033.073-3.844 1.222-4.53 1.38c-1.847.429-3.27 4.54-.772 3.646c5.563-1.99 4.704 4.56 7.732 5.136c.494.094 1.265.119 1.38.608c1.115 4.71 6.377 7.53 9.942 6.517c5.366-1.524 8.648 5.353 2.265 9.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.5 5.5h-25a6 6 0 0 0-6 6v25a6 6 0 0 0 6 6h25a6 6 0 0 0 6-6v-25a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}