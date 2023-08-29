package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GgOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><path fill="#e8112d" d="M192 0h128v512H192z"/><path fill="#e8112d" d="M0 187.7h512v136.6H0z"/><path id="flagGg1x10" fill="#f9dd16" d="m46 305.8l23.3-25h210v-49.7h-210L46 206.2z"/><use width="36" height="24" href="#flagGg1x10" transform="matrix(0 1.06667 -.9375 0 496 -17)"/><use width="36" height="24" href="#flagGg1x10" transform="matrix(0 -1.06667 .9375 0 16 529)"/><use width="36" height="24" href="#flagGg1x10" transform="rotate(180 256 256)"/>`),
		g.Group(children),
	)
}