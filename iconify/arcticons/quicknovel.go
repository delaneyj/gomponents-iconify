package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quicknovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 30.861a32.367 32.367 0 0 0-19.556 7.802A32.153 32.153 0 0 0 4.5 30.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.944 41.592v-26.92A34.983 34.983 0 0 1 43.5 6.407v27.37M23.944 14.672A34.775 34.775 0 0 0 4.5 6.407"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.944 41.592A32.335 32.335 0 0 1 43.5 33.776M4.5 6.407v27.37m19.444 7.816A32.112 32.112 0 0 0 4.5 33.776"/>`),
		g.Group(children),
	)
}