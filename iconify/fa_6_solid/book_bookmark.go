package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookBookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 96C0 43 43 0 96 0h96v190.7c0 13.4 15.5 20.9 26 12.5l54-43.2l54 43.2c10.5 8.4 26 .9 26-12.5V0h64c17.7 0 32 14.3 32 32v320c0 17.7-14.3 32-32 32v64c17.7 0 32 14.3 32 32s-14.3 32-32 32H96c-53 0-96-43-96-96V96zm64 320c0 17.7 14.3 32 32 32h256v-64H96c-17.7 0-32 14.3-32 32z"/>`),
		g.Group(children),
	)
}