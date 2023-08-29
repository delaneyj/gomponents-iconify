package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v3.413c.114.057.25.122.406.192c.61.27 1.547.62 2.81.9C8.482 6.786 10.076 7 12 7c1.925 0 3.52-.214 4.783-.495c1.264-.28 2.2-.63 2.81-.9c.158-.07.293-.135.407-.192V2h2v13h1v2h-1v5h-2v-5H4v5H2v-5H1v-2h1V2h2Zm0 13h2V8.267a16.767 16.767 0 0 1-2-.662V15Zm4-6.31V15h3V8.982a24.356 24.356 0 0 1-3-.291Zm5 .292V15h3V8.69c-.893.145-1.893.251-3 .292Zm5-.715V15h2V7.605c-.531.215-1.198.449-2 .662Zm2.446-3.1a.121.121 0 0 1 .003-.002l-.002.002h-.001Z"/>`),
		g.Group(children),
	)
}