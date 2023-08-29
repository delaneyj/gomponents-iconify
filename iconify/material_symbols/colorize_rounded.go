package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20v-3.325q0-.2.075-.388t.225-.337l8.65-8.65l-.75-.75q-.275-.275-.275-.7t.3-.725q.275-.275.7-.275t.725.275l1.175 1.2L16.9 3.25q.275-.275.7-.275t.7.275l2.4 2.4q.275.275.275.7t-.275.7l-3.075 3.075l1.25 1.25q.275.275.275.7t-.275.7q-.3.3-.725.3t-.7-.3l-.75-.725l-8.65 8.65q-.15.15-.338.225T7.325 21H4Zm1-2h1.95l8.3-8.35l-1.9-1.9L5 17.05V19Z"/>`),
		g.Group(children),
	)
}