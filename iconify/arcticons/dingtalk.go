package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dingtalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.36 28.8l-1.58 5.64h4.68l-2.7 8.84l9.87-12.5h-4.7l5.86-10.07a3.44 3.44 0 0 0-1.68-4.91l-27-11a1 1 0 0 0-1.38.89c-.16 5.73 3 7.94 4.07 8.33s13.51 5 13.51 5l-14.51-3.2a.5.5 0 0 0-.59.64c.63 2 2.54 6.89 6 7l9.12.29l-9.55 1.4a.5.5 0 0 0-.37.74c.89 1.54 3.43 5.11 7.24 4.11Z"/>`),
		g.Group(children),
	)
}