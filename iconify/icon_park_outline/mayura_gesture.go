package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MayuraGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20.995 36l-1.901-4.01l5.11-.417a3.782 3.782 0 0 0 0-7.565h-8.207C12.61 24.008 9 27.584 9 31.042c0 3.458.787 5.439 1.647 7.297c.859 1.858 3.15 5.661 9.355 5.661h8.953c5.248 0 9.045-4.855 9.032-11c-.008-3.83-.006-9.497.007-17a2.99 2.99 0 0 0-2.985-2.996h-.005a2.99 2.99 0 0 0-2.99 2.99V16c.01 5.141.015 8.158.015 9.051c0 2.894-1.34 3.891-4.022 2.992M25.955 24l.04-17a2.993 2.993 0 0 0-2.986-3h-.007a3 3 0 0 0-3 3v16.04"/><path d="M14.006 24V10a2.996 2.996 0 0 1 2.996-2.996h.004c1.657.003 3 1.347 3 3.005V24m6.001-3.98v-1a3 3 0 0 1 6 0v1"/></g>`),
		g.Group(children),
	)
}