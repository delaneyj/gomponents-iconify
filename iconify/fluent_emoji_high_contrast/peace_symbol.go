package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeaceSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M26 16c0 5.523-4.477 10-10 10S6 21.523 6 16S10.477 6 16 6s10 4.477 10 10Zm-11 7.938v-5.11l-3.669 3.67A7.956 7.956 0 0 0 15 23.937Zm-5.135-2.803L15 16V8.062a8.001 8.001 0 0 0-5.135 13.073ZM24 16a8.001 8.001 0 0 0-7-7.938V16l5.135 5.135A7.967 7.967 0 0 0 24 16Zm-7 2.828v5.11a7.956 7.956 0 0 0 3.669-1.44L17 18.827Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}