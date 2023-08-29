package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MujVlak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.541 26.622l2.713-15.534l7.583.003c7.024.002 11.708 5.785 10.462 12.917h0c-1.245 7.133-7.949 12.912-14.973 12.91l-7.907-.003l-7.381-.002C7.414 36.91 3.663 32.28 4.66 26.569h0c.997-5.71 6.365-10.338 11.989-10.337l2.636.001"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.61 26.623l1.814-10.386h2.514c4.225.002 7.042 3.48 6.293 7.77h0a9.573 9.573 0 0 1-9.006 7.764l-15.288-.005A4.217 4.217 0 0 1 9.73 26.57h0a6.4 6.4 0 0 1 6.021-5.192l2.637.001m1.797-10.293l-12.819-.004"/>`),
		g.Group(children),
	)
}