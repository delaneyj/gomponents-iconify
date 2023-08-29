package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416 208a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm208 112h-48V160A160 160 0 0 0 416 0H64A64 64 0 0 0 0 64v256a64 64 0 0 0 64 64h32a96 96 0 0 0 192 0h336a16 16 0 0 0 16-16v-32a16 16 0 0 0-16-16ZM192 432a48 48 0 1 1 48-48a48.05 48.05 0 0 1-48 48Zm64-240a32 32 0 0 1-32 32H96a32 32 0 0 1-32-32v-64a32 32 0 0 1 32-32h128a32 32 0 0 1 32 32Zm192 128H320V128a32 32 0 0 1 32-32h64a32 32 0 0 1 32 32Z"/>`),
		g.Group(children),
	)
}