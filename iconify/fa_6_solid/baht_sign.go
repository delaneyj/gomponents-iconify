package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BahtSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 0c-17.7 0-32 14.3-32 32v32H37.6C16.8 64 0 80.8 0 101.6v304.7c0 23 18.7 41.7 41.7 41.7H112v32c0 17.7 14.3 32 32 32s32-14.3 32-32v-32h32c61.9 0 112-50.1 112-112c0-40.1-21.1-75.3-52.7-95.1c13-18.3 20.7-40.7 20.7-64.9c0-61.9-50.1-112-112-112V32c0-17.7-14.3-32-32-32zm-32 128v96H64v-96h48zm64 96v-96c26.5 0 48 21.5 48 48s-21.5 48-48 48zm-64 64v96H64v-96h48zm64 96v-96h32c26.5 0 48 21.5 48 48s-21.5 48-48 48h-32z"/>`),
		g.Group(children),
	)
}