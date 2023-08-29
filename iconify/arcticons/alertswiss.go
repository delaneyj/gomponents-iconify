package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alertswiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.836 40.73v-7.155m0 11.925c7.065-2.715 6.66-8.154 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c0 3.246-.405 8.685 6.66 11.4Zm-3.578-8.347h7.155"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.54 42.743a21.503 21.503 0 1 1 10.959-18.97m.001.173a21.46 21.46 0 0 1-1.265 7.338"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.351 23.763a16.13 16.13 0 0 0-20.696.008M39.5 17.768a24.297 24.297 0 0 0-31 0m20.8 11.877a8.252 8.252 0 0 0-10.602-.03v.03"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.5 17.768l-15.501 18.05L8.5 17.768"/>`),
		g.Group(children),
	)
}