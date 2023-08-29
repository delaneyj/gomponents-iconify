package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M33.958 12.988C33.531 6.376 28.933 0 20.5 0C12.787 0 6.839 5.733 6.524 13.384C2.304 14.697 0 19.213 0 22.5C0 27.561 4.206 32 9 32h24c4.847 0 9-5.224 9-9.5c0-5.167-4.223-9.208-8.042-9.512zM33 31H9c-4.262 0-8-3.972-8-8.5C1 19.449 3.674 14 9 14h1.5a.5.5 0 0 0 0-1H9c-.509 0-.99.057-1.459.139C7.933 7.149 12.486 1 20.5 1C29.088 1 33 7.739 33 14v1.5a.5.5 0 0 0 1 0v-1.509c3.019.331 7 3.571 7 8.509c0 3.826-3.691 8.5-8 8.5z"/>`),
		g.Group(children),
	)
}