package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.875 22.65L1.35 4.125q-.275-.275-.275-.675t.3-.7q.275-.275.7-.275t.7.275L21.3 21.275q.275.275.275.675t-.3.7q-.275.275-.7.275t-.7-.275Zm-1.85-7.475L6.85 4h9.175q.825 0 1.413.588T18.025 6v4.5l3.15-3.15q.225-.225.538-.113t.312.463v8.6q0 .35-.312.463t-.538-.113l-3.15-3.15v1.675ZM4.025 4l14 14q-.025.8-.625 1.4T16 20H4q-.825 0-1.413-.588T2 18V6q0-.8.613-1.4T4.024 4Z"/>`),
		g.Group(children),
	)
}