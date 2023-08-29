package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2zm1.5 4l-.09.007a.5.5 0 0 0-.401.402L9 6.5v1.794L6.854 6.146l-.069-.058a.5.5 0 0 0-.568 0l-.07.058l-.058.069a.5.5 0 0 0 0 .568l.058.07L8.294 9H6.5l-.09.009a.5.5 0 0 0 0 .984L6.5 10h3l.058-.005l.05-.007l.064-.019l.04-.016l.076-.044l.065-.056l.044-.05l.036-.053l.02-.039l.023-.059l.013-.047l.01-.085V6.5l-.007-.09a.5.5 0 0 0-.402-.402L9.5 6l-.09.008L9.5 6z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}