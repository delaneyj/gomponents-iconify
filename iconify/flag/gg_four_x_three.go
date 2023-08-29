package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h640v480H0z"/><path fill="#e8112d" d="M256 0h128v480H256z"/><path fill="#e8112d" d="M0 176h640v128H0z"/><path id="flagGg4x30" fill="#f9dd16" d="m110 286.7l23.3-23.4h210v-46.6h-210L110 193.3z"/><use width="36" height="24" href="#flagGg4x30" transform="rotate(90 320 240)"/><use width="36" height="24" href="#flagGg4x30" transform="rotate(-90 320 240)"/><use width="36" height="24" href="#flagGg4x30" transform="rotate(180 320 240)"/>`),
		g.Group(children),
	)
}