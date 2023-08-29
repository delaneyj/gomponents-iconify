package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DarkMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.8 29.3c-8.9-.8-16.2-7.8-17.5-16.6c-.3-1.8-.3-3.7 0-5.4c.2-1.4-1.4-2.3-2.5-1.6C6.3 9.7 2.1 16.9 2.5 25c.5 10.7 9 19.5 19.7 20.4c10.6.9 19.8-6 22.5-15.6c.4-1.4-1-2.6-2.3-2c-2.9 1.3-6.1 1.8-9.6 1.5Z"/>`),
		g.Group(children),
	)
}