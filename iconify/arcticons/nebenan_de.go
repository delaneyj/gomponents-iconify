package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NebenanDe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5L7.556 13.722L24 21.944l16.444-8.222L24 5.5v37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.444 13.722v20.556l-4.11 2.055v-8.222l-8.223 4.111v8.222L24 42.5l-4.111-2.056v-8.222l-8.222-4.11v8.221l-4.111-2.055V13.722"/>`),
		g.Group(children),
	)
}