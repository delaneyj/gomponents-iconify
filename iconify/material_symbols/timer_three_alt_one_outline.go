package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThreeAltOneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.95 17.225h4.225q.875 0 1.5-.537T15.3 15.4v-1.3q0-.575-.463-.963t-1.112-.387q.65 0 1.113-.4t.462-.975v-.8q0-.75-.625-1.275t-1.5-.525H8.95v1.8h4.225v1.275h-2.1v1.8h2.1v1.75H8.95v1.825ZM9 3V1h6v2H9Zm3 19q-1.85 0-3.487-.713T5.65 19.35q-1.225-1.225-1.938-2.863T3 13q0-1.85.713-3.488T5.65 6.65q1.225-1.225 2.863-1.938T12 4q1.55 0 2.975.5t2.675 1.45l1.4-1.4l1.4 1.4l-1.4 1.4Q20 8.6 20.5 10.025T21 13q0 1.85-.713 3.488T18.35 19.35q-1.225 1.225-2.863 1.938T12 22Zm0-2q2.9 0 4.95-2.05T19 13q0-2.9-2.05-4.95T12 6Q9.1 6 7.05 8.05T5 13q0 2.9 2.05 4.95T12 20Z"/>`),
		g.Group(children),
	)
}