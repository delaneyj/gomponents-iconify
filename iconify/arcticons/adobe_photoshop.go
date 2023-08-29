package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobePhotoshop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.803 33V15h5.893c3.331 0 6.032 2.707 6.032 6.045s-2.7 6.045-6.032 6.045h-5.893m14.709 4.904c.822.69 1.709 1.006 3.7 1.006h1.01a2.978 2.978 0 0 0 2.975-2.981h0a2.978 2.978 0 0 0-2.974-2.981h-2.02a2.978 2.978 0 0 1-2.974-2.982h0a2.978 2.978 0 0 1 2.974-2.981h1.01c1.992 0 2.88.316 3.701 1.006"/>`),
		g.Group(children),
	)
}