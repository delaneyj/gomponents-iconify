package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 152a40 40 0 1 0-52.64 37.94A28 28 0 0 1 160 212h-24a28 28 0 0 1-28-28v-37.1c31.73-5.78 56-34.09 56-67.73V40a20 20 0 0 0-20-20h-16a12 12 0 0 0 0 24h12v35.17c0 24.4-19.47 44.52-43.41 44.83A44 44 0 0 1 52 80V44h12a12 12 0 0 0 0-24H48a20 20 0 0 0-20 20v40a68 68 0 0 0 56 66.93V184a52.06 52.06 0 0 0 52 52h24a52.06 52.06 0 0 0 51.61-45.72A40.08 40.08 0 0 0 240 152Zm-40 16a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}