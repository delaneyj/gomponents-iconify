package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onleihe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><rect width="11.98" height="15.87" x="10.17" y="18.61" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.99" transform="rotate(-20 16.163 26.562)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.5 26.86l-3.38-9.29A6 6 0 0 0 27.45 14h0a6 6 0 0 0-3.58 7.67L27.25 31m-3.38-9.33l-2.05-5.63"/>`),
		g.Group(children),
	)
}