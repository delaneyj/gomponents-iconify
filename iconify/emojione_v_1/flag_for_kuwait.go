package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForKuwait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ec1c24" d="M22.68 39L4.121 52.238C5.656 53.345 7.609 54 10 54h44c6.627 0 10-4.925 10-11v-4H22.68z"/><path fill="#137a08" d="M54 10H10c-2.201 0-4.03.553-5.514 1.5L22.879 25h41.12v-4c0-6.075-3.373-11-10-11"/><path fill="#25333a" d="M32.417 31.958L4.641 11.422C1.666 13.329 0 16.942 0 21v22c0 3.883 1.463 7.34 4.188 9.297L32.251 32.5l.166-.542z"/><path fill="#e6e7e8" d="M23 25h41v14H23z"/>`),
		g.Group(children),
	)
}