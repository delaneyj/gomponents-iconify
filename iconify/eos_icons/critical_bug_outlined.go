package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CriticalBugOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2v2h-2zm0-6h2v5h-2z"/><path fill="currentColor" d="M20 8h-2.81a5.985 5.985 0 0 0-1.82-1.96L17 4.41L15.59 3l-2.17 2.17A6.066 6.066 0 0 0 12 5a5.919 5.919 0 0 0-1.41.17L8.41 3L7 4.41l1.62 1.63A6.062 6.062 0 0 0 6.81 8H4v2h2.09A6.63 6.63 0 0 0 6 11v1H4v2h2v1a6.632 6.632 0 0 0 .09 1H4v2h2.81a5.99 5.99 0 0 0 10.38 0H20v-2h-2.09a6.632 6.632 0 0 0 .09-1v-1h2v-2h-2v-1a6.63 6.63 0 0 0-.09-1H20Zm-4 4v3a4.26 4.26 0 0 1-.07.7l-.1.65l-.37.65a3.993 3.993 0 0 1-6.92 0l-.37-.64l-.1-.65A4.271 4.271 0 0 1 8 15v-4a4.053 4.053 0 0 1 .07-.7l.1-.65l.37-.65a4.1 4.1 0 0 1 1.21-1.31l.57-.39l.74-.18A3.786 3.786 0 0 1 12 7a3.865 3.865 0 0 1 .95.12l.68.16l.61.42a3.894 3.894 0 0 1 1.21 1.31l.38.65l.1.65A4.038 4.038 0 0 1 16 11Z"/>`),
		g.Group(children),
	)
}