package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsSa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11s-4.935 11-11 11S5 22.065 5 16S9.935 5 16 5zm0 5c-2.206 0-4 1.794-4 4h-2l3 3l3-3h-2c0-1.103.897-2 2-2s2 .897 2 2v4c0 1.103-.897 2-2 2a1.995 1.995 0 0 1-1.723-1h-2.134c.447 1.721 1.998 3 3.857 3c2.206 0 4-1.794 4-4v-4c0-2.206-1.794-4-4-4z"/>`),
		g.Group(children),
	)
}