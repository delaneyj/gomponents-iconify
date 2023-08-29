package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covidalert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.04 11.85H18.96l-7.11 7.11v10.08l7.11 7.12h10.08l7.12-7.12V18.96l-7.12-7.11zM24 38.3l-3.58 6.2h7.16L24 38.3zm-10.11-4.19l-6.92 1.85l5.07 5.07l1.85-6.92zM9.7 24l-6.2-3.58v7.16L9.7 24zm4.19-10.11l-1.85-6.92l-5.07 5.07l6.92 1.85zM24 9.7l3.58-6.2h-7.16L24 9.7zm10.11 4.19l6.92-1.85l-5.07-5.07l-1.85 6.92zM38.3 24l6.2 3.58v-7.16L38.3 24zm-4.19 10.11l1.85 6.92l5.07-5.07l-6.92-1.85zM24 32.18v-4.13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 28.05l4.22.84l-.59-1.73l3.84-3.03l-1.02-.47l1.02-2.53l-2.37.38l-.59-1.39l-1.68 2.11h-.64l1.02-4.64l-1.73 1.01L24 15.82l-1.48 2.78l-1.73-1.01l1.02 4.64h-.64l-1.68-2.11l-.59 1.39l-2.37-.38l1.02 2.53l-1.02.47l3.84 3.03l-.59 1.73l4.22-.84z"/>`),
		g.Group(children),
	)
}