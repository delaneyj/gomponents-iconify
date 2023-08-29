package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quran(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.32 22.39h4v4h-4z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.32 26.39a2 2 0 0 1-2-2v6h-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2h-6v-6a2 2 0 1 1-2-2a2 2 0 0 1 2 2v-6h6a2 2 0 1 1 2-2a2 2 0 0 1-2 2h6v6a2 2 0 1 1 2 2ZM8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.73 4.5v39h24.92a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Z"/>`),
		g.Group(children),
	)
}