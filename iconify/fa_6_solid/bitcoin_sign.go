package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C48 14.3 62.3 0 80 0s32 14.3 32 32v32h32V32c0-17.7 14.3-32 32-32s32 14.3 32 32v32c0 1.5-.1 3.1-.3 4.5C254.1 82.2 288 125.1 288 176c0 24.2-7.7 46.6-20.7 64.9c31.7 19.8 52.7 55 52.7 95.1c0 61.9-50.1 112-112 112v32c0 17.7-14.3 32-32 32s-32-14.3-32-32v-32h-32v32c0 17.7-14.3 32-32 32s-32-14.3-32-32v-32h-6.3C18.7 448 0 429.3 0 406.3V101.6C0 80.8 16.8 64 37.6 64H48V32zm16 192h112c26.5 0 48-21.5 48-48s-21.5-48-48-48H64v96zm112 64H64v96h144c26.5 0 48-21.5 48-48s-21.5-48-48-48h-32z"/>`),
		g.Group(children),
	)
}