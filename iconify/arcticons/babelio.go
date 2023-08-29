package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Babelio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.737 43.5v-6.851a60.167 60.167 0 0 0 32.526-10.014V43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.917 29.702V15.279s-11.218 11.52-21.382 11.52v9.536"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.346 20.082V4.5s-3.54 9.938-10.842 12.122v9.366"/>`),
		g.Group(children),
	)
}