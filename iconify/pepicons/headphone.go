package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M17.88 15.069a1 1 0 0 1-1.898-.63a10.4 10.4 0 0 0 .518-3.273c0-4.456-2.756-7.412-6.5-7.412S3.5 6.71 3.5 11.166c0 1.14.178 2.247.518 3.273a1 1 0 0 1-1.898.63a12.4 12.4 0 0 1-.62-3.903c0-5.53 3.619-9.412 8.5-9.412s8.5 3.882 8.5 9.412c0 1.354-.212 2.673-.62 3.902Z"/><path d="M5.977 17.034a3 3 0 0 1-2.942-3.04v-.022a2.978 2.978 0 0 1 3.035-2.937a1 1 0 0 1 .98 1.013l-.054 4a1 1 0 0 1-1.019.986ZM14.089 11a3 3 0 0 1 2.942 3.04v.022A2.978 2.978 0 0 1 14.013 17h-.016a1 1 0 0 1-.981-1.014l.054-4A1 1 0 0 1 14.089 11Z"/></g>`),
		g.Group(children),
	)
}