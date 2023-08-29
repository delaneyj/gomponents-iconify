package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 4L7 6L5 8l2 2l2 2l.7-.7l.7-.7l-1.3-1.3L7.8 8l1.3-1.3l1.3-1.3l-.7-.7L9 4zm6 0l-.7.7l-.7.7l1.3 1.3L16.2 8l-1.3 1.3l-1.3 1.3l.7.7l.7.7l2-2l2-2l-2-2l-2-2zm5.977 15.09l-.388-1.532a.747.747 0 0 0-.356-.466a.764.764 0 0 0-.587-.072l-4.552 1.284L12 17l3.915-1.003a1.88 1.88 0 0 0 .558-.893L16.166 14L8 15H3v6h5l4.455.964a3.194 3.194 0 0 0 .727-.017l7.26-1.954a.755.755 0 0 0 .454-.344a.738.738 0 0 0 .081-.56ZM6 20H4v-4h2v4Zm6.922.982a.557.557 0 0 1-.138.018a.538.538 0 0 1-.115-.012L8 19.963V16.02l7.329-.908L10 17l4.721 2.232l.317.127l.328-.093l4.316-1.217l.264 1.042Z"/>`),
		g.Group(children),
	)
}