package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VeFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagVe4x30" transform="translate(0 -36)"><g id="flagVe4x31"><g id="flagVe4x32"><path id="flagVe4x33" fill="#fff" d="M0-5L-1.5-.2l2.8.9z"/><use width="180" height="120" href="#flagVe4x33" transform="scale(-1 1)"/></g><use width="180" height="120" href="#flagVe4x32" transform="rotate(72)"/></g><use width="180" height="120" href="#flagVe4x32" transform="rotate(-72)"/><use width="180" height="120" href="#flagVe4x31" transform="rotate(144)"/></g></defs><path fill="#cf142b" d="M0 0h640v480H0z"/><path fill="#00247d" d="M0 0h640v320H0z"/><path fill="#fc0" d="M0 0h640v160H0z"/><g id="flagVe4x34" transform="matrix(4 0 0 4 320 336)"><g id="flagVe4x35"><use width="180" height="120" href="#flagVe4x30" transform="rotate(10)"/><use width="180" height="120" href="#flagVe4x30" transform="rotate(30)"/></g><use width="180" height="120" href="#flagVe4x35" transform="rotate(40)"/></g><use width="180" height="120" href="#flagVe4x34" transform="rotate(-80 320 336)"/>`),
		g.Group(children),
	)
}