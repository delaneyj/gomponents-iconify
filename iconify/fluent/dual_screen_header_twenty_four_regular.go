package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenHeaderTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.748 4.002l-.001.002h7.498c.967 0 1.75.784 1.75 1.75V18.25a1.75 1.75 0 0 1-1.75 1.75h-8.997l-.001-.002H3.75A1.75 1.75 0 0 1 2 18.247V5.752c0-.967.784-1.75 1.75-1.75h8.998ZM20.495 7h-7.748v11.5h7.498a.25.25 0 0 0 .25-.25V7Zm-9.248 0H3.5v11.247c0 .138.112.25.25.25h7.498L11.247 7Z"/>`),
		g.Group(children),
	)
}