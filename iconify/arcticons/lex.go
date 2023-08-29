package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.957 17v13.913h6.956m9.043-1.653A3.366 3.366 0 0 1 24 31h0a3.489 3.489 0 0 1-3.478-3.478V25.26A3.489 3.489 0 0 1 24 21.783h0a3.489 3.489 0 0 1 3.478 3.478v1.217h-6.956m16.522-4.695L30.087 31m6.957 0l-6.957-9.217"/>`),
		g.Group(children),
	)
}