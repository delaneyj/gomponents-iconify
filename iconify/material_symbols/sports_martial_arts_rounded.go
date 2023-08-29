package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMartialArtsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.95 22q-.4 0-.687-.275t-.313-.675L9.5 13l-3.175-1.825l-.35 1.3L7.5 15.15q.2.35.088.75t-.463.6q-.35.2-.75.1t-.6-.45l-1.75-3.025q-.1-.175-.125-.375t.025-.4l.975-3.5q.05-.2.163-.35t.287-.25l5.4-3.1L8.7 3.1q-.275-.275-.288-.687T8.7 1.7q.275-.275.7-.275t.7.275l2.975 2.95q.35.35.288.838t-.488.737L10.4 7.65l1.2 1.05l7.5-6.125q.275-.25.663-.212t.687.387q.225.275.213.625T20.4 4l-7.9 8l-.45 9.05q-.025.4-.313.675T10.95 22ZM5 7q-.825 0-1.413-.588T3 5q0-.825.588-1.413T5 3q.825 0 1.413.588T7 5q0 .825-.588 1.413T5 7Z"/>`),
		g.Group(children),
	)
}