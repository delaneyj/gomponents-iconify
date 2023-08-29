package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 21.9q-.275.275-.7.275t-.7-.275L16 18.8V20q0 .825-.588 1.413T14 22h-4q-.825 0-1.413-.588T8 20v-9.2L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7ZM10 12.8V20h4v-3.2l-4-4Zm6 .35l-2-2v-.75l2-3V7H9.85l-2-2H16V4H6.85l-.725-.725q.225-.575.725-.925T8 2h8q.825 0 1.413.587T18 4v4l-2 3v2.15Zm-4 1.65Zm0-5.65Z"/>`),
		g.Group(children),
	)
}