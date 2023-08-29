package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.484.037c-1.494 0-2.791.786-3.469 1.947h6.936C11.275.823 9.977.037 8.484.037zm4.24 2.984H4.278c-.137 0-.248.102-.248.226v2.498c0 .124.111.225.248.225h8.446c.138 0 .248-.101.248-.225V3.247c0-.124-.111-.226-.248-.226zm-.742 3.985H5.009c.584 1.052 1.742 1.802 2.99 1.997v6.95h.262v.031h.597v-.031h.126V9.021a3.98 3.98 0 0 0 2.998-2.015z"/>`),
		g.Group(children),
	)
}