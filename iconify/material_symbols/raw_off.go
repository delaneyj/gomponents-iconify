package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RawOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM18.5 15l-.75-3.05l-.6 2.35l-2.45-2.4L14 9h1.5l.75 3L17 9h1.5l.75 3L20 9h1.5L20 15h-1.5Zm-9.75 0l1.05-4.2L11 12l1.5 1.5h-1.85l-.4 1.5h-1.5ZM3 15V9h3.5q.6 0 1.05.45T8 10.5v1q0 .45-.238.813T7.1 12.9L8 15H6.5l-.9-2H4.5v2H3Zm1.5-3.5h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}