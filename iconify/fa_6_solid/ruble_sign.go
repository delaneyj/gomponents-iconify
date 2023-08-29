package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 32c-17.7 0-32 14.3-32 32v192H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v32H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v32c0 17.7 14.3 32 32 32s32-14.3 32-32v-32h160c17.7 0 32-14.3 32-32s-14.3-32-32-32H128v-32h112c79.5 0 144-64.5 144-144S319.5 32 240 32H96zm144 224H128V96h112c44.2 0 80 35.8 80 80s-35.8 80-80 80z"/>`),
		g.Group(children),
	)
}