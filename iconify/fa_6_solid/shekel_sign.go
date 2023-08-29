package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShekelSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 32C14.3 32 0 46.3 0 64v384c0 17.7 14.3 32 32 32s32-14.3 32-32V96h128c35.3 0 64 28.7 64 64v160c0 17.7 14.3 32 32 32s32-14.3 32-32V160c0-70.7-57.3-128-128-128H32zm288 448c70.7 0 128-57.3 128-128V64c0-17.7-14.3-32-32-32s-32 14.3-32 32v288c0 35.3-28.7 64-64 64H192V192c0-17.7-14.3-32-32-32s-32 14.3-32 32v256c0 17.7 14.3 32 32 32h160z"/>`),
		g.Group(children),
	)
}