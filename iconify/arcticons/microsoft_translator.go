package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftTranslator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.315 35.636l4.142-11.144m3.97 11.177l-3.97-11.177m2.642 7.439h-5.407m14.166-18.584h12.457m-6.228-2.468v2.468m3.486 0c0 3.408-3.957 9.048-7.913 9.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.287 16.441c.392 2.35 4.466 6.268 8.09 6.816m-9.631 11.676a5.166 5.166 0 0 0 5.165-5.165v-2.576"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.581 29.522l2.33-2.33l2.331 2.33M19.537 13.347a5.166 5.166 0 0 0-5.165 5.165v2.575"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.701 18.758l-2.329 2.329l-2.332-2.329"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}