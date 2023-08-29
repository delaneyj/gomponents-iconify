package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PesoSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32c-17.7 0-32 14.3-32 32v64c-17.7 0-32 14.3-32 32s14.3 32 32 32v32c-17.7 0-32 14.3-32 32s14.3 32 32 32v160c0 17.7 14.3 32 32 32s32-14.3 32-32v-64h80c68.4 0 127.7-39 156.8-96H352c17.7 0 32-14.3 32-32s-14.3-32-32-32h-.7c.5-5.3.7-10.6.7-16s-.2-10.7-.7-16h.7c17.7 0 32-14.3 32-32s-14.3-32-32-32h-19.2C303.7 71 244.4 32 176 32H64zm190.4 96H96V96h80c30.5 0 58.2 12.2 78.4 32zM96 192h190.9c.7 5.2 1.1 10.6 1.1 16s-.4 10.8-1.1 16H96v-32zm158.4 96c-20.2 19.8-47.9 32-78.4 32H96v-32h158.4z"/>`),
		g.Group(children),
	)
}