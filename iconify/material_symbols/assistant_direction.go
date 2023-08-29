package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssistantDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23q-2.275 0-4.288-.863t-3.5-2.35q-1.487-1.487-2.35-3.5T1 12q0-2.3.863-4.3t2.35-3.488q1.487-1.487 3.5-2.35T12 1q2.3 0 4.3.863t3.488 2.35Q21.275 5.7 22.138 7.7T23 12q0 2.275-.863 4.288t-2.35 3.5q-1.487 1.487-3.487 2.35T12 23Zm-.625-3.225q.25.25.575.25t.575-.25l7.2-7.2q.25-.25.25-.6t-.25-.6l-7.2-7.2q-.25-.25-.575-.25t-.575.25l-7.2 7.2q-.25.25-.25.6t.25.6l7.2 7.2Zm-3.4-4.8v-4q0-.45.275-.725t.725-.275h4.15l-1.05-1.1l1.4-1.4l3.5 3.5l-3.5 3.5l-1.4-1.4l1.05-1.1h-3.15v3h-2Z"/>`),
		g.Group(children),
	)
}