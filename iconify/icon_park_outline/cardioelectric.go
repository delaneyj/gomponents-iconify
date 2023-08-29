package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cardioelectric(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5.555 23.194a18.75 18.75 0 0 1-.514-1.188C4.393 20.347 4 18.604 4 16.8C4 10.835 8.884 6 14.91 6c3.794 0 7.136 1.918 9.09 4.828A10.928 10.928 0 0 1 33.09 6C39.117 6 44 10.835 44 16.8c0 7.765-7.273 14.4-10.91 18c-2.423 2.4-5.454 4.8-9.09 7.2c-3.636-2.4-6.667-4.8-9.09-7.2c-.367-.362-.77-.756-1.199-1.178"/><path d="M8 29.973L19.114 19.04l6.464 6.666l9.684-9.887"/></g>`),
		g.Group(children),
	)
}