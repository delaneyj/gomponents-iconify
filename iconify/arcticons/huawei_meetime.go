package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiMeetime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="19.745" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="16.091" ry="15.245"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.724 12.187c2.539 2.72 7.338 7.616 1.836 12.472m-5.891 10.318c-1.006 2.83-2.572 6.815-6.419 8.523"/>`),
		g.Group(children),
	)
}