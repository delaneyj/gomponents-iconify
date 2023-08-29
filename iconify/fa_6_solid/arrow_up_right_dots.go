package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 0c-17.7 0-32 14.3-32 32s14.3 32 32 32h50.7L9.4 265.4c-12.5 12.5-12.5 32.8 0 45.3s32.8 12.5 45.3 0L256 109.3V160c0 17.7 14.3 32 32 32s32-14.3 32-32V32c0-17.7-14.3-32-32-32H160zm416 80a48 48 0 1 0-96 0a48 48 0 1 0 96 0zM448 208a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm-48 176a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm48 80a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm128 0a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm-304-80a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm48 80a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm-176 48a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm432-176a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm-48-80a48 48 0 1 0 0-96a48 48 0 1 0 0 96z"/>`),
		g.Group(children),
	)
}