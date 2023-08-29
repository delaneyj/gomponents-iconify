package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupervisorAccountRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 15q-1.05 0-1.775-.725T14.5 12.5q0-1.05.725-1.775T17 10q1.05 0 1.775.725T19.5 12.5q0 1.05-.725 1.775T17 15Zm-4 5q-.425 0-.713-.288T12 19v-.4q0-.6.313-1.113t.887-.737q.9-.375 1.863-.563T17 16q.975 0 1.938.188t1.862.562q.575.225.888.738T22 18.6v.4q0 .425-.288.713T21 20h-8Zm-3-8q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm-7 8q-.425 0-.713-.288T2 19v-1.8q0-.85.425-1.563T3.6 14.55q1.5-.75 3.112-1.15T10 13q.875 0 1.75.15t1.75.35l-1.7 1.7q-.625.625-1.213 1.275T10 18v.975q0 .3.113.563t.362.462H3Z"/>`),
		g.Group(children),
	)
}