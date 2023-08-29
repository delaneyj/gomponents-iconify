package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.1 18.15l-2-2V14h-2v.15l-2-2V12h4v-1q0-2.95-2.05-4.975T12.1 4q-1.1 0-2.087.313T8.2 5.2L6.75 3.8q1.125-.875 2.488-1.338T12.1 2q1.85 0 3.488.7t2.862 1.925q1.225 1.225 1.938 2.863T21.1 11v7.15Zm-.6 5.25l-.575-.575q-.15.125-.35.15T19.1 23h-7v-2h6l-1-1h-2v-2L5.575 8.475q-.225.6-.35 1.225T5.1 11v1h4v8h-4q-.825 0-1.412-.588T3.1 18v-7q0-1.125.238-2.113t.737-1.937L.7 3.575l1.425-1.4l19.775 19.8l-1.4 1.425ZM5.1 18h2v-4h-2v4Zm14-1.85ZM7.1 18h-2h2Z"/>`),
		g.Group(children),
	)
}