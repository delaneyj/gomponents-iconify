package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ftpclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v12.4c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2V7.5c0-1.1-.9-2-2-2Z"/><circle cx="13.6" cy="13.7" r="4.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 26.1h-33c-1.1 0-2 .9-2 2v12.4c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2V28.1c0-1.1-.9-2-2-2Z"/><circle cx="13.6" cy="34.3" r="4.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}