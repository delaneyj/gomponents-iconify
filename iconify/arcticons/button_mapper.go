package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonMapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2ZM5.6 24h36.8M24 42.4V5.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.7 28.6c2.6.5 4.3 2.9 3.8 5.4c-.3 1.9-1.9 3.5-3.8 3.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 31.7c.9.2 1.4 1 1.3 1.8c-.1.6-.6 1.2-1.3 1.3m-8.1-3.5v3.9H12l3.1 2.7v-9.4L12 31.3H8.9Zm25.5 3.3v2.9c0 .2.2.4.4.4h2c.2 0 .4-.2.4-.4h0v-3.7h.6c.4 0 .7-.3.7-.7c0-.2-.1-.4-.2-.5l-4.6-3.9c-.3-.3-.8-.3-1.1 0L28 32.6c-.3.2-.3.7-.1.9c.1.2.3.2.5.2h.6v3.7c0 .2.2.4.4.4h2c.2 0 .4-.2.4-.4v-2.9l2.6.1Zm-1.6-23.3l-4.1 4l4.1 4.1l-4.1-4.1h9.1M9.6 12.5h8.5v6.7H9.6z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 10.5H20v6.4"/>`),
		g.Group(children),
	)
}