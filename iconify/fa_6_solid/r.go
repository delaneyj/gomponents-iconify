package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func R(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v352c0 17.7 14.3 32 32 32s32-14.3 32-32V320h95.3l102.5 146.4c10.1 14.5 30.1 18 44.6 7.9s18-30.1 7.9-44.6l-84.2-120.2C282.8 288.1 320 236.4 320 176c0-79.5-64.5-144-144-144H64zm112 224H64V96h112c44.2 0 80 35.8 80 80s-35.8 80-80 80z"/>`),
		g.Group(children),
	)
}