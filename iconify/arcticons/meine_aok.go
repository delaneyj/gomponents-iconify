package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeineAok(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m28.767 26.868l-4.726 2.729m-.171-9.909v15.398m9.816.23h-19.64m9.644-5.846l-5.132-2.963"/><circle cx="24" cy="24" r="21.5"/><path d="m20.825 12.283l3.06-3.06h0l2.882 2.883m.109.12l.01.01a4.255 4.255 0 0 1 .015 6.021c-.003.003-.007.007-.015.005h0a4.26 4.26 0 0 1-6.025 0h0a4.26 4.26 0 0 1-.087-5.936M10.623 25.609l-1.12-4.18h0l3.938-1.054m.158-.035l.014-.004a4.26 4.26 0 0 1 5.218 3.013h0a4.26 4.26 0 0 1-8.197 2.325m23.623-5.309l4.138 1.265h0l-1.192 3.898m-.054.152a4.26 4.26 0 0 1-5.325 2.842h0a4.26 4.26 0 0 1 2.372-8.183"/></g>`),
		g.Group(children),
	)
}