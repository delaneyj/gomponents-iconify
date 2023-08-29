package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wordlesolver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 20.6v6.8h3.5M13 39.7h3.5M13 32.9h3.5M13 36.3h2.3M13 32.9v6.8m18 0v-6.8h2.2c1.3 0 2.3 1 2.3 2.3s-1 2.3-2.3 2.3H31m2.389-.007L35.5 39.7m0-19.1l-2.2 6.8l-2.3-6.8m-18.4-6.3c.4.5.9.8 1.7.8h1c.9 0 1.7-.8 1.7-1.7h0c0-.9-.8-1.7-1.7-1.7h-1.1c-.9 0-1.7-.8-1.7-1.7h0c0-.9.8-1.7 1.7-1.7h1c.8 0 1.3.2 1.7.8M31 12.8c0 1.3 1 2.3 2.2 2.3c1.3 0 2.3-1 2.3-2.3v-2.3c0-1.3-1-2.3-2.3-2.3s-2.2 1-2.2 2.3v2.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2ZM24 42.3V5.7M5.7 17.8h36.6M5.7 30.2h36.6"/>`),
		g.Group(children),
	)
}