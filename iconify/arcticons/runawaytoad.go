package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Runawaytoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.98" cy="25.85" r="6.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.98" cy="25.85" r="6.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="12.98" cy="25.85" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95" ry="4.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.57 26.84c-.47.37-.93.76-1.38 1.15a2.1 2.1 0 0 0-.69 1.57v8.63a2.12 2.12 0 0 0 2.11 2.11h34.78a2.12 2.12 0 0 0 2.11-2.11v-8.63a2.1 2.1 0 0 0-.69-1.56c-.45-.4-.91-.78-1.39-1.15m-11.24-5.32a28 28 0 0 0-6.18-.68a28.61 28.61 0 0 0-6.18.67"/><circle cx="35.02" cy="25.85" r="6.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.02" cy="25.85" r="6.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="35.02" cy="25.85" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95" ry="4.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.85 21.39L28.53 8.45l-2.93 6.22L24 7.7l-1.6 6.97l-2.93-6.22l-1.32 12.94m2.55 13.34l.8.81m5.8-.81l-.8.81"/>`),
		g.Group(children),
	)
}