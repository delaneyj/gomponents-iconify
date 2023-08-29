package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fitbod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M44.5 24.06L34.197 41.71l-20.5-.06L3.5 23.94L13.803 6.29l20.5.06L44.5 24.06Zm-41-.12l6.664.02"/><path d="m16.066 13.87l1.052-1.8l13.835.04l6.883 11.93l-6.954 11.89l-13.835-.04l-6.883-11.93l3.087-5.279m17.702-6.571l3.35-5.76"/><path d="M13.251 18.68h10.057l2.77-4.811H16.067"/></g>`),
		g.Group(children),
	)
}