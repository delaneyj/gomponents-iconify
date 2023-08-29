package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BtnMobileBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 42.5h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.945 20.55a2.694 2.694 0 0 1 0 5.387H10.5V15.161h4.445a2.694 2.694 0 0 1 0 5.388h0Zm0-.001H10.5m19.861 5.388V15.161L37.5 25.937V15.161m-17.153 0h7.139m-3.569 10.776V15.161M12.681 28.353H37.5l-27 4.486v-2.305c0-1.205.977-2.181 2.181-2.181Z"/>`),
		g.Group(children),
	)
}