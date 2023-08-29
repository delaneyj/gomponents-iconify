package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinalSurge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-26.141 9.872h8.627M12.838 24h5.607m-4.086-8.628l-3.043 17.256"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.56 30.737c.815 1.377 2.051 1.89 3.897 1.89h2.555a5.364 5.364 0 0 0 5.065-4.313h0A3.517 3.517 0 0 0 31.533 24H28.71a3.517 3.517 0 0 1-3.543-4.314h0a5.364 5.364 0 0 1 5.065-4.314h2.554c1.846 0 3.082.513 3.898 1.891"/>`),
		g.Group(children),
	)
}