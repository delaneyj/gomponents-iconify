package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redbubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 13.259V34.74a1 1 0 0 0 1 1h17.748a1 1 0 0 0 .787-1.616l-6.054-7.733l.097-.057c4.492-2.597 5.083-8.847 1.158-12.24h0a7.545 7.545 0 0 0-4.934-1.836H5.5a1 1 0 0 0-1 1Zm21.997 0V34.74a1 1 0 0 0 1 1h8.36a7.452 7.452 0 0 0 5.643-2.584l.19-.22a7.452 7.452 0 0 0-.517-10.278h0l.298-.724a7.007 7.007 0 0 0-3.413-8.97h0a7.007 7.007 0 0 0-3.065-.706h-7.496a1 1 0 0 0-1 1Z"/>`),
		g.Group(children),
	)
}