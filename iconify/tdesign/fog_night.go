package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FogNight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.916.585l-.45 1.675c-.3 1.116-.27 2.337.07 3.604a7 7 0 0 0 9.84 4.476l1.541-.752l-.103 1.71c-.063 1.034-.375 2.06-.92 3.15l-.446.894l-1.79-.895l.448-.894c.133-.266.245-.518.338-.76A9 9 0 0 1 8.604 6.38a9.59 9.59 0 0 1-.338-2.852a8.003 8.003 0 0 0-4.164 9.235c.017.061.085.251.174.48l.105.267l.033.082l.01.022l.002.007l.378.926l-1.852.756l-.378-.926l-.003-.01l-.01-.024l-.037-.09l-.113-.288a8.16 8.16 0 0 1-.24-.684C.74 7.947 3.906 2.464 9.24 1.034l1.675-.449ZM2.001 17h6v2h-6v-2Zm8 0h12v2h-12v-2Zm7 4h5v2h-5v-2Zm-15 0h13v2h-13v-2Z"/>`),
		g.Group(children),
	)
}