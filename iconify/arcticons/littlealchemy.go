package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Littlealchemy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.017 7.19l10.03 27.582a12.106 12.106 0 0 1 .729 4.137h0a4.592 4.592 0 0 1-1.345 3.247h0a4.591 4.591 0 0 1-3.247 1.344H15.816a4.591 4.591 0 0 1-3.247-1.345h0a4.592 4.592 0 0 1-1.345-3.247h0a12.106 12.106 0 0 1 .729-4.136L21.983 7.19h-1.345a1.345 1.345 0 0 1-1.345-1.345h0A1.345 1.345 0 0 1 20.638 4.5h6.724a1.345 1.345 0 0 1 1.345 1.345h0a1.345 1.345 0 0 1-1.345 1.345ZM13.1 31.617h21.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.522 27.565a1.278 1.278 0 1 0-.374-.904a1.294 1.294 0 0 0 .374.904Zm4.831-6.074a2.126 2.126 0 1 1 0 1.345a2.152 2.152 0 0 1 0-1.345Zm-2.37-3.543a.95.95 0 1 0-.279-.672a.962.962 0 0 0 .279.672Z"/>`),
		g.Group(children),
	)
}