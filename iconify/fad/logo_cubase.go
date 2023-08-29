package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoCubase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M58.446 127.855L140.401 46.3l32.106 26.083s-11.13.225-20.794 3.706c-9.663 3.48-14.486 8.596-19.58 14.04c-11.146 11.915-14.325 23.365-14.243 37.955c.08 14.59 4.462 22.448 12.94 33.133c2.87 3.619 5.696 7.099 10.508 9.728c1.64.897 8.563 5.947 13.122 7.36c4.559 1.413 16.789 1.949 16.789 1.949l-29.911 31.06l-82.892-83.459z"/><circle cx="173" cy="126" r="22"/></g>`),
		g.Group(children),
	)
}