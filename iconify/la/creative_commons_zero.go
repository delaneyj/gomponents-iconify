package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm0 5c-2.206 0-4 1.794-4 4v4c0 2.206 1.794 4 4 4s4-1.794 4-4v-4c0-2.206-1.794-4-4-4zm0 2c.256 0 .5.054.725.143L14 17.875V14c0-1.103.897-2 2-2zm2 2.12V18c0 1.103-.897 2-2 2c-.258 0-.502-.053-.729-.143L18 14.12z"/>`),
		g.Group(children),
	)
}