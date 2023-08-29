package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13.5 6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-5Z"/><path d="M4.5 3.029V28.5c0 1.115.921 2 2.034 2h19.432c.88 0 1.635-.55 1.917-1.33l.242-.67H24v-1h4V4a3 3 0 0 0-3-3H6.5c-1.114 0-2 .917-2 2.029ZM23 28.5H7.045a.512.512 0 0 1-.522-.5c0-.266.223-.5.522-.5H23v1ZM9 25V3h14v22H9Zm15 0V3h1a1 1 0 0 1 1 1v21h-2Z"/></g>`),
		g.Group(children),
	)
}