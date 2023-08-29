package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.424 4.266a.5.5 0 1 0-.848-.532l.848.532ZM7.35 12.062a.5.5 0 1 0 .847.532l-.847-.532ZM16.5 15a4.5 4.5 0 0 1-4.5 4.5v1a5.5 5.5 0 0 0 5.5-5.5h-1ZM12 19.5A4.5 4.5 0 0 1 7.5 15h-1a5.5 5.5 0 0 0 5.5 5.5v-1ZM7.5 15a4.5 4.5 0 0 1 4.5-4.5v-1A5.5 5.5 0 0 0 6.5 15h1Zm4.5-4.5a4.5 4.5 0 0 1 4.5 4.5h1A5.5 5.5 0 0 0 12 9.5v1Zm.576-6.766L7.35 12.062l.847.532l5.227-8.328l-.848-.532Z"/>`),
		g.Group(children),
	)
}