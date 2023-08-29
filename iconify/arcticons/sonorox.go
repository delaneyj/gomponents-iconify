package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sonorox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.62 7.6l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm-7.04 4.1l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm-7.04 4.1l-2.2-4.1H7.5l-2.2 4.1l2.2 4.1h4.84Zm0 8.2l-2.2-4.1H7.5L5.3 24l2.2 4.1h4.84Zm0 8.2l-2.2-4.1H7.5l-2.2 4.1l2.2 4.1h4.84Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.58 36.3l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm7.04 4.1l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm14.08-8.2l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.66 36.3l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm7.04-20.5l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm0 8.2l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Zm-7.04-12.3l-2.2-4.1h-4.84l-2.2 4.1l2.2 4.1h4.84Z"/>`),
		g.Group(children),
	)
}