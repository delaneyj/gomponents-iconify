package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkeeloAudiobooks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.994 18.407c1.983-3.405 5.612-4.527 10.027-2.956c-1.983 4.04-6.622 5.388-10.027 2.956Z"/><circle cx="16.079" cy="17.322" r="4.452" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.124 41.865a8.964 8.964 0 0 0 4.04-7.034c0-5.911-5.35-6.323-5.35-14.441c-1.87-1.61-6.532-9.551-4.152-14.03c.636-1.197 8.455 2.956 8.455 2.956c4.452-2.245 5.65-3.18 11.71-3.18S43 11.71 43 17.023c0 10.7-8.792 9.428-12.496 16.536a19.043 19.043 0 0 0-1.758 8.306Zm3.868-28.857l-.875-3.693"/>`),
		g.Group(children),
	)
}