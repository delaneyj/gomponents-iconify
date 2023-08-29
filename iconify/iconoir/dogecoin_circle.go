package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogecoinCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10 16.402V7.598c0-.331.268-.599.6-.604c2.49-.035 5.9-.072 5.9 5.006s-3.41 5.042-5.9 5.006a.606.606 0 0 1-.6-.604Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 12h4m0 10C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/></g>`),
		g.Group(children),
	)
}