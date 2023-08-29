package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Question(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.8 27.869a1 1 0 0 0-.8.98V31a1.5 1.5 0 0 1-3 0v-4.322a1.5 1.5 0 0 1 1.395-1.496c1.778-.124 4.393-.73 6.515-2.083c2.071-1.322 3.59-3.285 3.59-6.254c0-2.86-1.061-4.843-2.602-6.158c-1.578-1.347-3.774-2.085-6.132-2.177c-4.881-.191-9.43 2.31-10.302 6.204a1.5 1.5 0 0 1-2.928-.655c.666-2.974 2.587-5.151 5.074-6.576c2.5-1.433 5.531-2.078 8.274-1.97c2.868.112 5.76 1.013 7.962 2.892c2.24 1.912 3.654 4.744 3.654 8.44c0 4.238-2.27 7.056-4.976 8.783c-1.832 1.169-3.9 1.87-5.723 2.24ZM21 40a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}