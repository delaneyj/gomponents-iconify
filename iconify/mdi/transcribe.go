package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transcribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7c0-1.11.89-2 2-2h16m-2 12v-2h-5.5l-2 2H18M6 17h2.5l6.85-6.88c.2-.19.2-.51 0-.71l-1.76-1.76c-.2-.2-.52-.2-.71 0L6 14.53V17Z"/>`),
		g.Group(children),
	)
}