package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TbNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsTbNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm16.525 12c5.439 0 5.365 6.872 5.327 10.4l-.004.377c.05-.024.1-.05.148-.078c.44-.257 1.004-.801 1.004-2.199V6h2v14.5c0 1.398.564 1.942 1.004 2.199c.049.028.098.054.148.078l-.004-.377c-.038-3.528-.112-10.4 5.327-10.4c7.45 0 12.773 26.007 9.58 28.66c-3.194 2.654-9.58 1.062-12.773-2.123c-2.578-2.57-2.355-8.578-2.195-12.891c.01-.263.02-.52.028-.769a4.273 4.273 0 0 1-1.119-.45a3.816 3.816 0 0 1-.996-.835c-.297.352-.64.626-.996.834c-.4.234-.794.37-1.12.451l.029.767v.002c.16 4.313.383 10.32-2.195 12.891c-3.193 3.185-9.58 4.777-12.773 2.123C3.752 38.007 9.075 12 16.525 12ZM32 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm2 11a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm3-1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-3-6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTbNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}