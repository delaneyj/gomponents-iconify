package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BobWorld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.59 29.161l-1.781 7.01l-1.78-7.01l-1.78 7.01l-1.781-7.01m10.119 7.009h0a1.772 1.772 0 0 1-1.78-1.752V33.28a1.772 1.772 0 0 1 1.78-1.752h0a1.772 1.772 0 0 1 1.78 1.752v1.14a1.772 1.772 0 0 1-1.78 1.752Zm3.789-2.89a1.772 1.772 0 0 1 1.78-1.753h0m-1.78 0v4.644m3.625-7.01v7.01m5.531-2.891a1.772 1.772 0 0 0-1.78-1.753h0a1.772 1.772 0 0 0-1.78 1.752v1.14a1.772 1.772 0 0 0 1.78 1.752h0a1.772 1.772 0 0 0 1.78-1.753m0 1.753v-7.01M25 26.517h0a3.638 3.638 0 0 1-3.581-3.672v-2.387A3.638 3.638 0 0 1 25 16.786h0a3.638 3.638 0 0 1 3.581 3.672v2.387A3.638 3.638 0 0 1 25 26.517Zm-14.002-6.059a3.638 3.638 0 0 1 3.58-3.672h0a3.638 3.638 0 0 1 3.582 3.672v2.387a3.638 3.638 0 0 1-3.581 3.672h0a3.638 3.638 0 0 1-3.581-3.672m0 3.672V11.829m20.842 8.629a3.638 3.638 0 0 1 3.581-3.672h0a3.638 3.638 0 0 1 3.581 3.672v2.387a3.638 3.638 0 0 1-3.58 3.672h0a3.638 3.638 0 0 1-3.582-3.672m0 3.672V11.829"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}