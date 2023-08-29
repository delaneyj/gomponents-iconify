package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeepassTwoAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.15 20.67a3.07 3.07 0 1 1 3.06 3.07a3.07 3.07 0 0 1-3.06-3.07Zm18.64 3.07a3.07 3.07 0 1 1 3.07-3.07a3.07 3.07 0 0 1-3.07 3.07ZM24 30.47a2.61 2.61 0 0 1 1.59 4.69l1.24 4.31h-5.67l1.23-4.31A2.62 2.62 0 0 1 24 30.47Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.51 28.31V40a1.94 1.94 0 0 0 1.94 2h33.1a2 2 0 0 0 2-2V28.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39 13.72l2.58-2.56a3 3 0 0 0-4.27-4.26h0l-2.55 2.6a18.39 18.39 0 0 0-21.51 0l-2.57-2.57A3.22 3.22 0 0 0 8.5 6a3 3 0 0 0-2.12 5.15L9 13.71a18.41 18.41 0 0 0-3.47 10.8v3.8h37v-3.8A18.57 18.57 0 0 0 39 13.72Zm-19.87 7a2.93 2.93 0 1 1-2.93-2.92a2.93 2.93 0 0 1 2.94 2.9Zm12.65 2.92a2.92 2.92 0 1 1 2.92-2.92a2.92 2.92 0 0 1-2.91 2.9ZM5.51 34.97h13.77m9.44 0H42.5"/>`),
		g.Group(children),
	)
}