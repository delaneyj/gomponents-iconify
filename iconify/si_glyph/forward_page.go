package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.798 3.516L11.892.266c-.277-.288-.593-.335-.871-.048c0 0 .037 2.19.006 2.19c-.68 0-1.083.113-1.7.314A6.329 6.329 0 0 0 7.119 3.99C5.843 5.113 5.026 6.392 5.026 8.203c0 1.182.523.752.742.339c1.035-1.945 2.923-2.928 5.281-2.928c.03 0-.001 2.146-.001 2.146c.278.287.628.253.869.047l3.881-3.25a.754.754 0 0 0 0-1.041z"/><path d="M14 8.079v4.983H.938V3.937h5.215c.276-.312.621-.717 1.223-.919H.008v10.964H14.97V7.258l-.97.821z"/></g>`),
		g.Group(children),
	)
}