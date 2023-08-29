package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lyfpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="35.764" cy="28.537" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 14.408v14.343h7.171m8.891-3.586v4.841a3.586 3.586 0 0 1-3.586 3.586h0a3.574 3.574 0 0 1-2.535-1.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.062 19.249v5.916a3.586 3.586 0 0 1-3.586 3.586h0a3.586 3.586 0 0 1-3.585-3.586V19.25m12.064 9.501V16.918a2.51 2.51 0 0 1 2.51-2.51h0A2.922 2.922 0 0 1 37 15.458m-7.171 3.791h5.02"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}