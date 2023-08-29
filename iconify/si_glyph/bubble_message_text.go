package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.01.078C3.606.078.031 2.759.031 6.065c0 3.018 2.98 5.508 6.852 5.92l-2.121 3.438l6.197-3.799c2.945-.88 5.031-3.035 5.031-5.559c0-3.306-3.574-5.987-7.98-5.987zM10 9H4V8h6v1zm2-1.954H4V6h8v1.046zm0-2.055H4v-1h8v1z"/>`),
		g.Group(children),
	)
}