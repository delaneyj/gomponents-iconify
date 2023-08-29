package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm12.983 5.482l.603.798a7 7 0 0 1-.003 8.44l-.603.798l-1.595-1.206l.603-.798a5 5 0 0 0 .002-6.029l-.603-.798l1.596-1.205ZM8 16a4 4 0 0 0-4 4h9.05v2H2v-2a6 6 0 0 1 6-6h5v2H8Zm8.19-.91l.603.799a3.5 3.5 0 0 1-.001 4.22l-.604.798l-1.595-1.206l.604-.798a1.5 1.5 0 0 0 0-1.809l-.602-.798l1.595-1.205Z"/>`),
		g.Group(children),
	)
}