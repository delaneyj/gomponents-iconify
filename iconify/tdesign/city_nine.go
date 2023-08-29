package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.01 1.828l3.988 2.492l3.99-2.492l6.384 3.829l-1.029 1.715L21 7.166V22H2.997V7.166l-.343.206l-1.029-1.715l6.384-3.83ZM4.996 5.967V20H7v-5h2v5h1.997V6.053l-3.01-1.88l-2.99 1.794ZM13 20h2v-5h2v5h2V5.967l-2.991-1.794l-3.01 1.88V20ZM7 8.998h2.004v2.004H7V8.998Zm8 0h2.004v2.004H15V8.998Z"/>`),
		g.Group(children),
	)
}