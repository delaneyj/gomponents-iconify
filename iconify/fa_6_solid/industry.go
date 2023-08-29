package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Industry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32c-17.7 0-32 14.3-32 32v368c0 26.5 21.5 48 48 48h416c26.5 0 48-21.5 48-48V152.2c0-18.2-19.4-29.7-35.4-21.1L352 215.4v-63.2c0-18.2-19.4-29.7-35.4-21.1L160 215.4V64c0-17.7-14.3-32-32-32H64z"/>`),
		g.Group(children),
	)
}