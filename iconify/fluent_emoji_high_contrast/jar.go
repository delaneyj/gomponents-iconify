package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M23 8a1 1 0 0 0-1 1v18a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1ZM9.999 14H21v7H9.999A1 1 0 0 1 9 20v-5a1 1 0 0 1 .999-1Z"/><path d="M8 2.5A1.5 1.5 0 0 1 9.5 1h13a1.5 1.5 0 0 1 0 3H22v1h2.105C26.203 5 28 6.622 28 8.738v18.524C28 29.378 26.203 31 24.105 31H7.895C5.797 31 4 29.378 4 27.262V8.738C4 6.622 5.797 5 7.895 5H10V4h-.5A1.5 1.5 0 0 1 8 2.5ZM7.895 7C6.795 7 6 7.83 6 8.738v18.524C6 28.17 6.795 29 7.895 29h16.21c1.1 0 1.895-.83 1.895-1.738V8.738C26 7.83 25.205 7 24.105 7H7.895Z"/></g>`),
		g.Group(children),
	)
}