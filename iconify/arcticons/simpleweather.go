package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simpleweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.43 17.33c4.46-.24 9.1 2.33 10.64 8.21a6.88 6.88 0 0 1 .05 13.71H11.89c-9.21-1.25-10.47-14.66 0-16.46a10.14 10.14 0 0 1 8.54-5.46Zm0 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.55 18.37a6.22 6.22 0 1 1 8.58 8.21m-3-14.3V8.75m-6.25 6.12l-2.5-2.5m15 15.01l2.5 2.49m.09-8.75h3.53m-6.12-6.25l2.5-2.5"/>`),
		g.Group(children),
	)
}