package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h2v-7q0-.825.213-1.625T6.85 6.85l1.5 1.5q-.175.4-.262.813T8 10v7h6.2L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4l-3.65-3.6H4Zm14-3.85l-2-2V10q0-1.65-1.175-2.825T12 6q-.65 0-1.25.2t-1.1.6L8.2 5.35q.5-.4 1.075-.7T10.5 4.2V2h3v2.2q2 .5 3.25 2.113T18 10v5.15Zm-6.9-1.25ZM12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm.825-12.025Z"/>`),
		g.Group(children),
	)
}