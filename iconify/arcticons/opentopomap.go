package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opentopomap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2.003 2.003 0 0 0-2 2v33a2.003 2.003 0 0 0 2 2h33a2.003 2.003 0 0 0 2-2v-33a2.003 2.003 0 0 0-2-2Zm30.2 19.1l4.78-.736M5.5 29.6l18.4-2.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.1 8.38a10.2 10.2 0 0 1 12.496 7.206l.004.014a10.2 10.2 0 0 1-6.45 12.2l.144.52l1.11.645l2.83 10.6l-3.32.892l-2.83-10.6l.641-1.11l-.14-.526a10.2 10.2 0 0 1-4.47-19.9Zm.819 3.09a6.908 6.908 0 1 0 .003 0ZM33.1 27.8l-1.6.5m-13.6-.6L14.2 5.5"/>`),
		g.Group(children),
	)
}