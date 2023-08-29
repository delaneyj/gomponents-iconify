package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pharmacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M.5 6.5V10c0 4.56 2.5 8.5 7 10.36v.14l-.123.123a12 12 0 0 1-3.119 2.248l-.758.379v.25h17v-.25l-.758-.38a12 12 0 0 1-3.119-2.247L16.5 20.5v-.14c4.5-1.86 7-5.8 7-10.36V6.5H.5Z"/><path d="M10.5 12.5v-3h3v3h3v3h-3v3h-3v-3h-3v-3h3Zm9-12l-6 6h4l4-4V2A1.5 1.5 0 0 0 20 .5h-.5Z"/></g>`),
		g.Group(children),
	)
}