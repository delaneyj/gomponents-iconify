package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iriunwebcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.464 33.868a8.687 8.687 0 1 1 8.687-8.687a8.696 8.696 0 0 1-8.687 8.687Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.396 36.486A11.417 11.417 0 1 0 13.979 25.07a11.43 11.43 0 0 0 11.417 11.417M24 4.5a17.219 17.219 0 1 0 17.219 17.219A17.218 17.218 0 0 0 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.81 37.537a38.817 38.817 0 0 1 6.883 5.346a3.306 3.306 0 0 1-.558.45a49.878 49.878 0 0 0-12.464-1.4a53.822 53.822 0 0 0-13.39 1.567a14.58 14.58 0 0 1-.98-.637c1.045-1.055 5.852-4.848 7.102-5.236"/>`),
		g.Group(children),
	)
}