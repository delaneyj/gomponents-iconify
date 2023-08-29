package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prtg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.31 34.99c.86-1.8 1.34-3.81 1.34-5.93c0-7.61-6.16-13.77-13.76-13.77c-7.61 0-13.77 6.16-13.77 13.77c0 3.03.98 5.84 2.64 8.11L7.7 39.23a19.398 19.398 0 0 1-3.57-11.24c0-10.78 8.74-19.53 19.53-19.53c10.78 0 19.53 8.75 19.53 19.53c0 4.36-1.43 8.39-3.85 11.64l-6.03-4.64Zm4.079-20.884l-5.726 6.388m-21.579-6.54l3.082 3.702"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.364 22.075l16.991 9.589l-2.86 2.579l-14.131-12.168z"/>`),
		g.Group(children),
	)
}