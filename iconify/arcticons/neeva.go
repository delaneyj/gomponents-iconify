package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neeva(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m4.5 28.51l.002 12.232h12.08V28.51H4.5m26.918-2.76l.002 14.992H43.5V25.75M4.653 25.7c.542-10.68 9.615-18.915 20.299-18.42"/><path d="M25.22 18.271a7.376 7.376 0 0 1 6.204 7.568M28.662 7.813a19.426 19.426 0 0 1 14.8 18.018m-38.793-.078a19.427 19.427 0 0 0 20.51-18.444m3.384.695s.271 5.93-3.294 10.22"/></g>`),
		g.Group(children),
	)
}