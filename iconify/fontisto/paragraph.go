package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.545 0H6.181a6.182 6.182 0 0 0 0 12.364h2.546V22.94a1.091 1.091 0 0 0 2.182 0v-.032v.002V2.183h2.182v20.76a1.091 1.091 0 0 0 2.182 0v-.032v.002V2.184h3.304a1.091 1.091 0 0 0 0-2.182h-.032h.002zM6.182 10.182a4 4 0 0 1 0-8h2.546v8z"/>`),
		g.Group(children),
	)
}