package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swatchbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 32C0 14.3 14.3 0 32 0h128c17.7 0 32 14.3 32 32v384c0 53-43 96-96 96S0 469 0 416V32zm223.6 393.9c.3-3.3.4-6.6.4-9.9V154l75.4-75.4c12.5-12.5 32.8-12.5 45.3 0l90.5 90.5c12.5 12.5 12.5 32.8 0 45.3L223.6 425.9zM182.8 512l192-192H480c17.7 0 32 14.3 32 32v128c0 17.7-14.3 32-32 32H182.8zM128 64H64v64h64V64zM64 192v64h64v-64H64zm32 248a24 24 0 1 0 0-48a24 24 0 1 0 0 48z"/>`),
		g.Group(children),
	)
}