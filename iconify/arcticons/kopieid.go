package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kopieid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.5v-25a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v25a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.051 17.5h13v13h-13zm18.074 3.796h8.252m-8.281 5.467l12.197.028"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.055 28.43a9.468 9.468 0 0 0-4.516-2.372a3.168 3.168 0 1 0-3.975 0A9.466 9.466 0 0 0 8.05 28.43"/>`),
		g.Group(children),
	)
}