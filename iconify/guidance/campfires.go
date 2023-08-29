package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campfires(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 17L9 19.402m6 1.696l8.5 2.402m-23 0l23-6.5M12 17c-3.314 0-6-2.15-6-5.425v-.242a5.11 5.11 0 0 1 1.524-3.636L10.02 5.23C11.288 3.977 12 1.772 12 0c0 1.772.712 3.977 1.98 5.23l2.496 2.466A5.113 5.113 0 0 1 18 11.334v.242C18 14.85 15.313 17 12 17Zm0 0c-1.381 0-2.5-.914-2.5-2.262c0-.561.229-1.17.635-1.567l1.04-1.015c.528-.516.825-1.423.825-2.152c0 .73.297 1.636.825 2.152l1.04 1.015c.407.397.635 1.006.635 1.567C14.5 16.086 13.38 17 12 17Z"/>`),
		g.Group(children),
	)
}