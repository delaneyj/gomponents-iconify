package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Join(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="19.953" height="35.053" x="21.593" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.798"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.593 9.994h19.953M21.593 36.059h19.953M21.593 7.447H8.252a1.798 1.798 0 0 0-1.798 1.798v31.457A1.798 1.798 0 0 0 8.252 42.5H24.61a1.798 1.798 0 0 0 1.798-1.798M6.454 11.941h15.139M6.454 38.006h15.139"/><circle cx="15.483" cy="25.187" r="1.895" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.378 25.187h4.215m0-1.895h9.711m-3.082 3.347l3.347-3.347l-3.347-3.347"/>`),
		g.Group(children),
	)
}