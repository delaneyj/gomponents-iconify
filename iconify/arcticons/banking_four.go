package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.154 36.57l5.305-5.305l-5.305-5.304"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.579 19.88H39.42a3.072 3.072 0 0 1 3.079 3.067V39.57a3.072 3.072 0 0 1-3.066 3.079H8.58a3.072 3.072 0 0 1-3.08-3.065V22.96a3.072 3.072 0 0 1 3.066-3.08h.013Zm33.921-.801v-3.183a3.072 3.072 0 0 0-3.067-3.078H8.58a3.072 3.072 0 0 0-3.08 3.065v3.197m37-7.167V8.65a3 3 0 0 0-3.001-3H8.58A3.072 3.072 0 0 0 5.5 8.715v3.197"/>`),
		g.Group(children),
	)
}