package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxTissue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m383.88 287.82l64-192H338.47a70.2 70.2 0 0 1-66.59-48a70.21 70.21 0 0 0-66.6-48H63.88l64 288Zm-384 192a32 32 0 0 0 32 32h448a32 32 0 0 0 32-32v-64h-512Zm480-256h-40.94l-21.33 64h14.27a16 16 0 0 1 0 32h-352a16 16 0 1 1 0-32h15.21l-14.22-64h-49a32 32 0 0 0-32 32v128h512v-128a32 32 0 0 0-31.99-32Z"/>`),
		g.Group(children),
	)
}