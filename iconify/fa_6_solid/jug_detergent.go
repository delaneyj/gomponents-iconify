package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JugDetergent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 24c0-13.3 10.7-24 24-24h80c13.3 0 24 10.7 24 24v24h8c13.3 0 24 10.7 24 24s-10.7 24-24 24H88c-13.3 0-24-10.7-24-24s10.7-24 24-24h8V24zM0 256c0-70.7 57.3-128 128-128h128c70.7 0 128 57.3 128 128v192c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V256zm256 0v96c0 17.7 14.3 32 32 32s32-14.3 32-32v-96c0-17.7-14.3-32-32-32s-32 14.3-32 32z"/>`),
		g.Group(children),
	)
}