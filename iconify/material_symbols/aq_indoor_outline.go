package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AqIndoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.55 13.075l.7-1.925h.1l.675 1.925H8.55Zm5.875 1.375q-.725 0-1.25-.537T12.65 12.6q0-.775.525-1.312t1.25-.538q.725 0 1.25.538T16.2 12.6q0 .375-.125.713t-.35.587l-.575-.85l-.575.4l.55.85q-.175.075-.337.113t-.363.037ZM7 15.05h.85l.45-1.25h2l.425 1.25h.875l-1.875-4.925H8.85L7 15.05Zm8.925.4l.575-.4l-.35-.525q.425-.375.638-.875T17 12.6q0-1.075-.75-1.838T14.425 10q-1.05 0-1.8.763t-.75 1.837q0 1.075.75 1.825t1.8.75q.275 0 .563-.075t.562-.2l.375.55ZM4 20V8l8-6l8 6v12H4Zm2-2h12V9l-6-4.5L6 9v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}