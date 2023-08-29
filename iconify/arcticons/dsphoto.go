package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dsphoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5a2 2 0 0 1 2 2h0v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2h0v-33a2 2 0 0 1 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.7 12.1a4.5 4.5 0 1 0 4.5 4.5a4.55 4.55 0 0 0-4.5-4.5ZM28 21l-7.2 7.2a.67.67 0 0 1-1 0h0l-1.4-1.4a.67.67 0 0 0-1 0l-7.8 7.8a.67.67 0 0 0 0 1a.76.76 0 0 0 .5.2h27.8a.68.68 0 0 0 .7-.7a.6.6 0 0 0-.1-.4L29 21.1a.76.76 0 0 0-1-.1Z"/>`),
		g.Group(children),
	)
}