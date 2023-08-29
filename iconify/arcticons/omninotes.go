package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Omninotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.46 11.54l-5.85-5.9a3.91 3.91 0 0 0-2.76-1.14h-19.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V14.3a3.9 3.9 0 0 0-1.19-2.76Zm-25.08 3.15H29.7m-16.32 7.03h21.24m-21.24 7.03h19.5m-19.5 7.03h12.68"/>`),
		g.Group(children),
	)
}