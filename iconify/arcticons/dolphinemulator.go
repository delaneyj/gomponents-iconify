package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolphinemulator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.46 34.36c-4-10.1-7.9-13.5-11.6-16.4c-1.8-1.4-.5-3.2 2.7-4.1c-2.2-.3-4.9-.7-8.3 1.7c-3.9-1.8-18.7-4.1-20.3 6.4l-1.2 1.6c-.6 1.9.1 1.7 1.4 1s8-2.5 9.9-1.5c1.5 3.5 3.5 4.1 8.3 4.8c-1.4-.8-3.4-1.4-3.9-4.4c3.1-.1 10.3 2.3 13.3 5.2c-1-2.4-10.8-10-23.2-7c-.4-2.3 4.3-5.1 9.8-4.7c5.3.4 10.4 1.8 14.1 5.1a34.4 34.4 0 0 1 9 12.3Z"/>`),
		g.Group(children),
	)
}