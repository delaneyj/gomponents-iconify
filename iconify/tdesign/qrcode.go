package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qrcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2V2Zm2 2v5h5V4H4Zm9-2h9v9h-9V2Zm2 2v5h5V4h-5ZM5.5 5.5h2.004v2.004H5.5V5.5Zm11 0h2.004v2.004H16.5V5.5Zm-3.504 7.496H15V15h-2.004v-2.004Zm7 0H22V15h-2.004v-2.004ZM2 13h9v9H2v-9Zm2 2v5h5v-5H4Zm11.996.996H18v2h2v2h2V22h-2.004v-2h-2v-2h-2v-2.004ZM5.5 16.5h2.004v2.004H5.5V16.5Zm7.496 3.496H15V22h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}