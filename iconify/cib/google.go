package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.318 13.714v5.484h9.078c-.37 2.354-2.745 6.901-9.078 6.901c-5.458 0-9.917-4.521-9.917-10.099s4.458-10.099 9.917-10.099c3.109 0 5.193 1.318 6.38 2.464l4.339-4.182C24.251 1.584 20.641.001 16.318.001c-8.844 0-16 7.151-16 16s7.156 16 16 16c9.234 0 15.365-6.49 15.365-15.635c0-1.052-.115-1.854-.255-2.651z"/>`),
		g.Group(children),
	)
}