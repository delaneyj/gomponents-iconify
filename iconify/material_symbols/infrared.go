package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infrared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.9 21.9l-1.425-1.425q1.65-1.65 2.588-3.813T18 12q0-2.5-.938-4.663t-2.587-3.812L15.9 2.1q1.9 1.9 3 4.438T20 12q0 2.925-1.1 5.463t-3 4.437Zm-2.825-2.825L11.65 17.65q1.1-1.1 1.725-2.55T14 12q0-1.65-.625-3.1T11.65 6.35l1.425-1.425q1.35 1.35 2.138 3.175T16 12q0 2.075-.788 3.9t-2.137 3.175ZM10.25 16.25l-1.425-1.425q.55-.55.863-1.275T10 12q0-.825-.313-1.55t-.862-1.275L10.25 7.75q.8.8 1.275 1.9T12 12q0 1.25-.475 2.35t-1.275 1.9ZM4 20V4h2v6q.825 0 1.413.588T8 12q0 .825-.588 1.413T6 14v6H4Z"/>`),
		g.Group(children),
	)
}