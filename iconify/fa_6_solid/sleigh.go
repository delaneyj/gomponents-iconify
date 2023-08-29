package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 32C14.3 32 0 46.3 0 64s14.3 32 32 32v160c0 53 43 96 96 96v32h64v-32h192v32h64v-32c53 0 96-43 96-96v-96c17.7 0 32-14.3 32-32s-14.3-32-32-32h-64c-17.7 0-32 14.3-32 32v41.3c0 30.2-24.5 54.7-54.7 54.7c-75.5 0-145.6-38.9-185.6-102.9l-4.3-6.9C174.2 67.6 125 37.6 70.7 32.7c-2.2-.5-4.4-.7-6.7-.7H32zm608 352c0-17.7-14.3-32-32-32s-32 14.3-32 32v8c0 13.3-10.7 24-24 24H64c-17.7 0-32 14.3-32 32s14.3 32 32 32h488c48.6 0 88-39.4 88-88v-8z"/>`),
		g.Group(children),
	)
}