package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mizuni(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm-4.5 5c-.83 0-1.5.67-1.5 1.5V22c.79-.78 1.82-1.41 3-1.8v-8.7c0-.83-.67-1.5-1.5-1.5zm4.5 0c-.83 0-1.5.67-1.5 1.5v8.33c.49-.08.99-.12 1.51-.12c.51 0 1 .04 1.49.12V11.5c0-.83-.67-1.5-1.5-1.5zm4.5 0c-.83 0-1.5.67-1.5 1.5v8.69c1.18.4 2.21 1.02 3 1.8V11.5c0-.83-.67-1.5-1.5-1.5z"/>`),
		g.Group(children),
	)
}