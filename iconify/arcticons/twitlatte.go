package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitlatte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.66 14.91h-1.38v-4a.61.61 0 0 0-.61-.6H10.28a.6.6 0 0 0-.6.6V29.6a5.48 5.48 0 0 0 1.88 4.12h-5a.7.7 0 0 0-.66.73a3.24 3.24 0 0 0 3.27 3.19h28.7a3.23 3.23 0 0 0 3.23-3.23a.69.69 0 0 0-.69-.69h-5a5.51 5.51 0 0 0 1.89-4.08v-2.32h1.11a6.39 6.39 0 0 0 6.44-5.6v-.63a6.18 6.18 0 0 0-6.19-6.18Zm0 8.57h-1.38v-4.75h1.38a2.43 2.43 0 0 1 0 4.75Zm-27.1 10.24h23.83"/>`),
		g.Group(children),
	)
}