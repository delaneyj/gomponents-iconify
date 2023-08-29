package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.34 17c2.76 4.783 8.876 6.42 13.66 3.66a9.956 9.956 0 0 0 4.196-4.731a9.985 9.985 0 0 0-.536-8.93a9.985 9.985 0 0 0-7.465-4.928A9.956 9.956 0 0 0 7 3.339C2.217 6.101.578 12.217 3.34 17Z"/><path stroke-linecap="round" d="M16.95 20.573S16.01 13.983 14 10.5c-2.01-3.482-6.95-7.073-6.95-7.073m.527 17.39c1.482-4.47 8.875-9.424 14.286-8.237m-5.45-9.371C14.927 7.63 7.675 12.512 2.29 11.452" opacity=".5"/></g>`),
		g.Group(children),
	)
}