package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meijer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 23.518a1.91 1.91 0 0 1 1.905-1.905h0a1.91 1.91 0 0 1 1.905 1.905v3.143M9.5 21.613v5.143m3.81-3.238a1.91 1.91 0 0 1 1.905-1.905h0a1.91 1.91 0 0 1 1.905 1.905v3.143m5.592-.953c-.286.572-.953.953-1.62.953h0a1.91 1.91 0 0 1-1.904-1.905v-1.238a1.91 1.91 0 0 1 1.904-1.905h0a1.91 1.91 0 0 1 1.905 1.905v.666h-3.81"/><circle cx="25.131" cy="19.517" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.131 21.613v5.143m3.239-5.143v5.715a1.91 1.91 0 0 1-1.906 1.905h0a1.88 1.88 0 0 1-1.333-.572"/><circle cx="28.369" cy="19.517" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.141 25.708c-.286.572-.952.953-1.62.953h0a1.91 1.91 0 0 1-1.904-1.905v-1.238a1.91 1.91 0 0 1 1.905-1.905h0a1.91 1.91 0 0 1 1.905 1.905v.666h-3.81m5.978-.666a1.91 1.91 0 0 1 1.905-1.905h0m-1.905 0v5.143"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}