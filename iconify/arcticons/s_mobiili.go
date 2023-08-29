package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SMobiili(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.863 43.5h2.534a15.885 15.885 0 0 0 15.885-15.885v-5.821a4.538 4.538 0 0 1 4.538-4.539h.806v-6.967h-1.414A10.212 10.212 0 0 0 24 20.5v7a10.212 10.212 0 0 1-10.212 10.212h-1.414v-6.967h.806a4.538 4.538 0 0 0 4.538-4.539v-5.821A15.885 15.885 0 0 1 33.603 4.5h0a15.885 15.885 0 0 0-15.885 15.885v5.821a4.538 4.538 0 0 1-4.538 4.539h-.806v6.967h1.415A10.212 10.212 0 0 0 24 27.5"/>`),
		g.Group(children),
	)
}