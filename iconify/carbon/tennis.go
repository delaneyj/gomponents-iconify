package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<rect width="4" height="4" x="20" y="2" fill="currentColor" rx="2"/><path fill="currentColor" d="M11 16a1 1 0 0 1-.707-1.707l7-7a1 1 0 1 1 1.414 1.414l-7 7A.996.996 0 0 1 11 16Z"/><path fill="currentColor" d="M23.707 24.293a8.395 8.395 0 0 0-4.72-2.209c.2-.164.393-.336.578-.521a9.245 9.245 0 0 0 1.733-9.914l-1.877.697c1.11 2.98.635 5.898-1.272 7.802a6.696 6.696 0 0 1-5.56 1.805a9.775 9.775 0 0 1-5.772-2.786c-2.973-2.97-4.076-8.227-.98-11.32c1.906-1.905 4.826-2.38 7.809-1.27l.699-1.875c-3.683-1.369-7.486-.706-9.924 1.731c-3.943 3.939-2.676 10.496.98 14.149a11.814 11.814 0 0 0 6.976 3.36c.21.022.416.023.623.033V24h4.996a6.846 6.846 0 0 1 4.297 1.707L26.586 30L28 28.586Z"/>`),
		g.Group(children),
	)
}