package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NbaTwoKTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.151 22.65c.289-1.635 1.998-2.925 3.634-2.6c1.074.214 1.792 1.175 1.704 2.302c-.065.837-.476 1.664-1.144 2.172c-1.238.94-5.137 3.476-5.137 3.476h5.3M10.36 22.65c.288-1.635 1.997-2.925 3.634-2.6c1.074.214 1.792 1.175 1.704 2.302c-.066.837-.477 1.664-1.145 2.172C13.315 25.464 9.416 28 9.416 28h5.3m19.807 0h0c-1.464 0-2.44-1.186-2.183-2.65l.476-2.7c.259-1.464 1.654-2.65 3.118-2.65h0c1.463 0 2.44 1.186 2.182 2.65l-.476 2.7c-.258 1.464-1.653 2.65-3.117 2.65Zm-15.531-8l-1.411 8m.492-2.787l5.214-5.186M21.881 28l-2.588-4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}