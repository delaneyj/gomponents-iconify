package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M9 12a3 3 0 0 0 5.121 2.121l3.36 3.36A10.44 10.44 0 0 1 12 19c-4.664 0-8.4-2.903-10-7c.901-2.307 2.48-4.236 4.52-5.48l3.359 3.359A2.99 2.99 0 0 0 9 12Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.73 5.073A10.96 10.96 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.595 11.595 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}