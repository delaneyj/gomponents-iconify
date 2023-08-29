package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nemid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.789 23.905h-19.22m-2.072 2.074v2.979m21.292-5.053a4.86 4.86 0 1 1 1.032 2.998M13.34 24.379v4.375m-2.771-4.849a1.892 1.892 0 0 0-2.071 2.074"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}