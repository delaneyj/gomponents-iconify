package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vials(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 64c0-17.7 14.3-32 32-32h160c17.7 0 32 14.3 32 32s-14.3 32-32 32v304c0 44.2-35.8 80-80 80s-80-35.8-80-80V96C14.3 96 0 81.7 0 64zm136 32H88v160h48V96zm152-32c0-17.7 14.3-32 32-32h160c17.7 0 32 14.3 32 32s-14.3 32-32 32v304c0 44.2-35.8 80-80 80s-80-35.8-80-80V96c-17.7 0-32-14.3-32-32zm136 32h-48v160h48V96z"/>`),
		g.Group(children),
	)
}