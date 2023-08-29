package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.196 22.261c-1.186-.175-2.265.456-2.265 2.202v.413a2.696 2.696 0 0 0 2.697 2.697h0a2.696 2.696 0 0 0 2.696-2.697v-1.752a2.696 2.696 0 0 0-2.696-2.697c-1.282 0-1.801.272-2.516 1.05m-7.883-4.69v10.786m0-4.449a2.696 2.696 0 0 1 2.696-2.697h0a2.696 2.696 0 0 1 2.696 2.697v4.449m-12.285-.603c.492.413 1.024.603 2.218.603h.605c.984 0 1.782-.8 1.782-1.787h0c0-.986-.798-1.786-1.782-1.786h-1.21c-.985 0-1.783-.8-1.783-1.786h0c0-.987.798-1.787 1.783-1.787h.605c1.193 0 1.725.19 2.217.603m17.669 3.846a2.696 2.696 0 0 0 2.697 2.697h0a2.696 2.696 0 0 0 2.697-2.697v-1.752a2.696 2.696 0 0 0-2.697-2.697h0a2.696 2.696 0 0 0-2.696 2.697m0-2.697v10.786"/>`),
		g.Group(children),
	)
}