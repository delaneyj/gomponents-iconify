package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindMiniFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17.035a.5.5 0 0 1-.788.408l-7.133-5.035a.5.5 0 0 1 0-.817l7.133-5.035a.5.5 0 0 1 .788.409v10.07Zm1.079-4.627a.5.5 0 0 1 0-.817l7.133-5.035a.5.5 0 0 1 .788.409v10.07a.5.5 0 0 1-.788.408l-7.133-5.035Z"/>`),
		g.Group(children),
	)
}