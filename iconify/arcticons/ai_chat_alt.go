package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AiChatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.705 27.607l-5.07-2.928a12.807 12.807 0 0 1-6.404-11.091V7.73a5.23 5.23 0 0 0-5.23-5.23h0a5.23 5.23 0 0 0-5.23 5.23v5.856c0 4.576-2.44 8.804-6.403 11.092l-5.073 2.928a5.23 5.23 0 0 0-1.914 7.144h0a5.23 5.23 0 0 0 7.144 1.914l5.072-2.928a12.807 12.807 0 0 1 12.807 0l5.072 2.928a5.23 5.23 0 0 0 7.143-1.914h0a5.23 5.23 0 0 0-1.914-7.144Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.295 20.393l5.07 2.928a12.807 12.807 0 0 1 6.404 11.091v5.858a5.23 5.23 0 0 0 5.23 5.23h0a5.23 5.23 0 0 0 5.23-5.23v-5.856c0-4.576 2.44-8.804 6.403-11.092l5.073-2.928a5.23 5.23 0 0 0 1.914-7.144h0a5.23 5.23 0 0 0-7.144-1.914l-5.072 2.928a12.807 12.807 0 0 1-12.807 0l-5.072-2.928a5.23 5.23 0 0 0-7.143 1.914h0a5.23 5.23 0 0 0 1.914 7.144Z"/><circle cx="24" cy="24" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}