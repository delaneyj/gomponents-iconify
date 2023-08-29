package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldikoNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.483 8.539s4.864-.092 6.391 4.753c.703 2.23.127 6.766-4.409 8.251c-2.556.838-9.242.43-9.242-6.626c0-2.415.15-4.767 5.319-7.328c5.041-2.499 7.8-2.107 10.466-2.063c4.528.074 15.542 1.53 15.542 10.777c0 6.007.017 16.548.162 22.553c.02.819.091 2.265 2.478 2.143c0 0-10.903 5.819-15.878-5.628c0 0-1.807 6.963-8.482 6.963c-7.24 0-10.962-1.735-12.19-7.785c-.855-4.21 1.51-8.822 7.255-10.336l7.24-1.456s6.523-1.477 6.391-9.99c0 0 .17-8.019-7.362-6.797c0 0-2.878.249-3.68 2.57v-.001Z"/><path d="M22.045 40.922c-5.144-3.56-4.6-15.051-.481-18.726"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.312 35.371V15.449"/>`),
		g.Group(children),
	)
}