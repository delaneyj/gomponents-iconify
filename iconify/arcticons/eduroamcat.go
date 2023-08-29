package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eduroamcat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.07 29.56A19.61 19.61 0 0 1 5.75 9h0m36.62.4h0A19.54 19.54 0 0 1 38 29.56m-14-19c1.36 0 2.46 2 2.46 4.43s-1.1 4.42-2.46 4.42s-2.46-2-2.46-4.42s1.1-4.4 2.46-4.4Zm.08 12.52v15.94M15.62 9.45c-1.14 1.23-1.86 4.46-1.86 6.05c0 6.83 4.4 11.74 10.32 11.74m8.47-17.79c1.14 1.23 1.86 4.46 1.86 6.05c0 6.83-4.4 11.74-10.33 11.74"/>`),
		g.Group(children),
	)
}