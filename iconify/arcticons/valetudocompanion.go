package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Valetudocompanion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.02 6.482a18.498 18.498 0 0 1-4.13 23.42c-3.437 2.885-10.012 6.07-11.89 11.615c-2.064-5.51-8.451-8.729-11.891-11.615a18.498 18.498 0 0 1-4.13-23.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.409 17.974c0 4.644-4.526 7.01-8.409 9.53c-3.769-2.409-8.409-4.886-8.409-9.53a8.409 8.409 0 0 1 16.818 0Z"/><circle cx="24" cy="17.974" r="2.803" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}