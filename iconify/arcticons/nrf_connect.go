package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NrfConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.712 10.803c5.4-10.482 17.114-.999 15.665 5.918c-1.253 5.978-3.785 12.004-9.59 16.633c-5.065 4.04-19.144.13-11.125-14.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.234 36.738c-7.122 11.738-18.053.587-16.603-6.33c1.252-5.978 3.785-12.004 9.59-16.633c5.065-4.04 18.647 1.734 10.844 15.852"/>`),
		g.Group(children),
	)
}