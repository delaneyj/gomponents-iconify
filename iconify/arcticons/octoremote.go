package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octoremote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.872 27a5.983 5.983 0 0 0-.81 2.055a5.405 5.405 0 0 0 1.12 4.669c1.687 1.866 4.82 2.254 7.423 1.72c2.067-.424 4.009-1.98 5.52-3.394c1.464-1.37 2.525-2.606 2.905-2.546c.773.121-2.95 12.391-13.335 13.088S9.559 38.565 8.981 28.096C8.283 15.464 22.615 4.678 22.615 4.678S34.742 3.15 39.043 9.051c-.32 1.505-2.757 3.85-5.74 6.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.61 5.895c-.417.257-3.222 4.585-7.342 7.876c-.704.562-1.44 1.132-2.185 1.724m-3.125 2.67a18.942 18.942 0 0 0-5.739 9.356a9.12 9.12 0 0 0 8.618 11.639a14.185 14.185 0 0 0 12.687-9.36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.022 33.456V17.31a1.613 1.613 0 0 1 1.447-1.745h10.233a1.613 1.613 0 0 1 1.447 1.745v14.7m0 4.47v5.275a1.613 1.613 0 0 1-1.447 1.745H23.47a1.394 1.394 0 0 1-1.195-.758"/>`),
		g.Group(children),
	)
}