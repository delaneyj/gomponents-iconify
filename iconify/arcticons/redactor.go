package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redactor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V20.41l-2.86 5L28.22 16L38.67 5.5ZM28.22 16l11.42 9.42s-3.06-.54-10.83 5.48l-4.25-3.7s4.5-8.74 3.66-11.2Zm-6.87 12.64h3.21l-2.48 3.3l-.73.54a3.32 3.32 0 0 1-1.3.24a2.18 2.18 0 0 1-1.2-.3a1.57 1.57 0 0 1-.48-.48H17a1.1 1.1 0 0 1-.3.35a2 2 0 0 1-1.27.36h-1.75v-.71H9.81l-.22.71H8.51v-4h1.41l.26-.74h1.23l.26.74h2v-.72h1.47a3.14 3.14 0 0 1 1.43.26a.9.9 0 0 1 .4.46Zm4-.77L27.83 30l-1.49 1.93h-4.26Z"/>`),
		g.Group(children),
	)
}