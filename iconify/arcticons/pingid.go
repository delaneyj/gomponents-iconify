package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pingid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.856 8.431v8.11H4.5v-8.11h8.356m4.371 30.897V8.43h9.642c3.35 0 6.61-.012 9.641 1.656c8.785 4.835 9.422 21.145 1.157 26.79c-5.602 3.828-14.031 2.45-20.44 2.45m8.356-24.716v18.795c2.21 0 4.512.019 6.294-1.496c3.069-2.608 3.088-8.284 2.235-11.893c-.307-1.3-.88-2.641-1.856-3.582c-1.815-1.751-4.304-1.824-6.673-1.824m-12.727 6.05v18.667H4.5V20.66h8.356Z"/>`),
		g.Group(children),
	)
}