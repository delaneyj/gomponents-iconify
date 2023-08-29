package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lyft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.397 25.246v5.083a3.777 3.777 0 0 1-3.765 3.766h0a3.16 3.16 0 0 1-2.636-1.13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.397 19.032v6.214a3.777 3.777 0 0 1-3.765 3.765h0a3.777 3.777 0 0 1-3.766-3.765v-6.214m18.74-3.203v13.18m-2.07-9.98h3.953M11.514 13.9v13.18a1.78 1.78 0 0 0 1.882 1.886h.567M27.445 29V16.574a2.574 2.574 0 0 1 2.636-2.636h0a2.932 2.932 0 0 1 2.636 1.13m-7.343 3.954h5.272"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}