package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wlanscanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 18.4c10.5-10.3 28.3-10.3 39-.1m-32.1 7a18.51 18.51 0 0 1 25.2-.2M19 32.5h10l-5.1 4.8Z"/>`),
		g.Group(children),
	)
}