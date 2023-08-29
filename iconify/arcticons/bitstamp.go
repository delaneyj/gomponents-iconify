package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitstamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.467 21.32a8.41 8.41 0 0 1 0 16.818H14.59V4.5h13.876a8.41 8.41 0 0 1 0 16.82Zm0-.001H14.591m0-16.819h-3.467m3.467 16.819h-3.467m3.467 16.819h-3.467m0 5.362h25.752"/>`),
		g.Group(children),
	)
}