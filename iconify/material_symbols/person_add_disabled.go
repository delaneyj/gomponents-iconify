package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAddDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L17 19.8v.2H1v-2.8q0-.85.438-1.563T2.6 14.55q1.55-.775 3.15-1.163T9 13q.325 0 .638.013t.612.037L9.2 12H9q-1.65 0-2.825-1.175T5 8v-.2L1.4 4.2l1.425-1.425l18.4 18.4L19.8 22.6ZM18 14v-3h-3V9h3V6h2v3h3v2h-3v3h-2Zm-5.4-4.3L7.3 4.4q.375-.2.813-.3T9 4q1.65 0 2.825 1.175T13 8q0 .45-.1.888t-.3.812Z"/>`),
		g.Group(children),
	)
}