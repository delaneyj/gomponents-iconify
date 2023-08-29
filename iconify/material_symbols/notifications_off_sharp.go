package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h2v-7q0-.825.213-1.625T6.85 6.85L10 10H7.2L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4l-3.65-3.6H4Zm14-3.85l-9.8-9.8q.5-.4 1.075-.7T10.5 4.2V2h3v2.2q2 .5 3.25 2.113T18 10v5.15ZM12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Z"/>`),
		g.Group(children),
	)
}