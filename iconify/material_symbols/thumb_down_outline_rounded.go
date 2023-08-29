package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16q-.8 0-1.4-.6T1 14v-2q0-.175.05-.375t.1-.375l3-7.05q.225-.5.75-.85T6 3h11v13l-6 5.95q-.375.375-.888.438t-.987-.188q-.475-.25-.7-.7t-.1-.925L9.45 16H3Zm12-.85V5H6l-3 7v2h9l-1.35 5.5L15 15.15ZM20 3q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-3v-2h3V5h-3V3h3Zm-5 2v10.15V5Z"/>`),
		g.Group(children),
	)
}