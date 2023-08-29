package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandBackFist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 0c-26.5 0-48 21.5-48 48v128c0 8.8-7.2 16-16 16s-16-7.2-16-16v-26.7l-9 7.5C40.4 169 32 187 32 206v38c0 38 16.9 74 46.1 98.3L128 384v96c0 17.7 14.3 32 32 32h160c17.7 0 32-14.3 32-32V374.7c46.9-19 80-65 80-118.7V144c0-26.5-21.5-48-48-48c-12.4 0-23.6 4.7-32.1 12.3C350 83.5 329.3 64 304 64c-12.4 0-23.6 4.7-32.1 12.3C270 51.5 249.3 32 224 32c-12.4 0-23.6 4.7-32.1 12.3C190 19.5 169.3 0 144 0z"/>`),
		g.Group(children),
	)
}