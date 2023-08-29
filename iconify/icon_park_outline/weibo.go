package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weibo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4.159" d="M12.562 16.41c-3.344 3.25-10.399 12.736-5.2 19.522c5.2 6.785 19.92 4.168 25.997-.613c6.077-4.78 5.477-7.734 4.16-9.036c-1.318-1.301-5.163.014-6.24-1.757c-1.076-1.77 1.917-6.195-.634-7.595c-2.55-1.4-6.954 3.65-8.944 2.338c-1.989-1.312 2.19-5.596 0-6.83c-2.189-1.235-5.795.722-9.14 3.972Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4.159" d="M43.379 16.516a12.954 12.954 0 0 0-3.642-8.073a12.963 12.963 0 0 0-7.723-3.84m5.278 12.328c-.13-1.64-.867-3.117-2-4.228a7.249 7.249 0 0 0-4.24-2.01"/><path fill="currentColor" d="M25 30c0 2.21-2.686 4-6 4s-6-1.79-6-4s2.686-4 6-4s6 1.79 6 4Z"/></g>`),
		g.Group(children),
	)
}