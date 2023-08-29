package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CtrlCProgrammingIdleGame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.054 24.27l-.02-11.095H39.26V43.5H17.206V24.252Zm-.021-11.095L17.206 24.252"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.206 34.825H8.74V15.577l10.848.018L19.566 4.5h11.228v8.675M19.566 4.5L8.739 15.577"/>`),
		g.Group(children),
	)
}