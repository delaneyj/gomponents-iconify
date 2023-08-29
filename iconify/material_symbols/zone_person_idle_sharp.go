package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZonePersonIdleSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 22.05l-7.275-7.3l-.4 1.95h-4l1-5.1l-1.8.7v3.4H6V11l2.4-1.025L.675 2.25L2.1.825l19.8 19.8l-1.425 1.425ZM2 20.7v-5h2v3h3v2H2Zm18-15v-3h-3v-2h5v5h-2Zm-18 0v-3h2v3H2Zm15 15v-2h3v2h-3ZM6.825 2.7l-2-2H7v2h-.175ZM22 17.875l-2-2V15.7h2v2.175ZM13.5 8.2q-.825 0-1.413-.587T11.5 6.2q0-.825.588-1.413T13.5 4.2q.825 0 1.413.588T15.5 6.2q0 .825-.588 1.413T13.5 8.2Z"/>`),
		g.Group(children),
	)
}