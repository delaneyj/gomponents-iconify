package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PemutarVisha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsPemutarVisha0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.576 21.857v-3.322a1.374 1.374 0 0 1 2.071-1.18l9.403 5.564a1.372 1.372 0 0 1-.023 2.374l-2.358 1.339l-9.093-4.775Zm9.093 4.775h0l-7.058 4.034a1.37 1.37 0 0 1-2.034-1.201a1.37 1.37 0 0 0 2.049 1.161l7.043-3.994Zm-9.092-4.775h0l4.004 2.102h-.002l-4.002-2.102Zm4.002 2.102l5.09 2.673l-7.058 4.034a1.37 1.37 0 0 1-2.034-1.201v-3.403l4.002-2.103Zm.002 0l-4.005 2.103v-4.206l4.005 2.103Zm.933 18.535c-.171.004-.342.006-.514.006c-10.21 0-18.5-8.29-18.5-18.5S13.79 5.5 24 5.5S42.5 13.79 42.5 24c0 .172-.002.343-.006.514v16.919c0 .582-.481 1.06-1.064 1.06H24.514ZM24 10.576c-7.41 0-13.424 6.015-13.424 13.424S16.591 37.424 24 37.424S37.424 31.409 37.424 24S31.409 10.576 24 10.576Z"/></defs><use href="#arcticonsPemutarVisha0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsPemutarVisha0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsPemutarVisha0" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}