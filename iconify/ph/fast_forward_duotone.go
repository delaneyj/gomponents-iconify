package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M136 128a7.76 7.76 0 0 1-3.63 6.59l-88.19 56.15A7.91 7.91 0 0 1 32 184.15V71.85a7.91 7.91 0 0 1 12.18-6.59l88.19 56.15A7.76 7.76 0 0 1 136 128Zm100.37-6.59l-88.19-56.15A7.91 7.91 0 0 0 136 71.85v112.3a7.91 7.91 0 0 0 12.18 6.59l88.19-56.15a7.8 7.8 0 0 0 0-13.18Z" opacity=".2"/><path d="M240.67 114.66L152.48 58.5A15.91 15.91 0 0 0 128 71.84v37.3L48.48 58.5A15.91 15.91 0 0 0 24 71.84v112.32a15.92 15.92 0 0 0 24.48 13.34L128 146.86v37.3a15.92 15.92 0 0 0 24.48 13.34l88.19-56.16a15.8 15.8 0 0 0 0-26.68ZM40 183.94V72.07L127.82 128Zm104 0V72.07L231.82 128Z"/></g>`),
		g.Group(children),
	)
}