package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroneDelivery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 30h-8v-8h8Zm-6-2h4v-4h-4Z"/><path fill="currentColor" d="M32 11H22v2h4v3H6v-3h4v-2H0v2h4v5h5.132L6 22.697V27h2v-3.697L11.535 18h8.93L24 23.303V27h2v-4.303L22.868 18H28v-5h4v-2zM16 6a5.982 5.982 0 0 0-4.24 1.76l1.413 1.413a3.994 3.994 0 0 1 5.654 0l1.414-1.414A5.981 5.981 0 0 0 16 6z"/><path fill="currentColor" d="m8.932 4.932l1.414 1.414a7.988 7.988 0 0 1 11.308 0l1.414-1.414a9.984 9.984 0 0 0-14.136 0Z"/>`),
		g.Group(children),
	)
}