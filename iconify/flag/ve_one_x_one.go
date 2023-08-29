package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VeOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagVe1x10" transform="translate(0 -36)"><g id="flagVe1x11"><g id="flagVe1x12"><path id="flagVe1x13" fill="#fff" d="M0-5L-1.5-.2l2.8.9z"/><use width="180" height="120" href="#flagVe1x13" transform="scale(-1 1)"/></g><use width="180" height="120" href="#flagVe1x12" transform="rotate(72)"/></g><use width="180" height="120" href="#flagVe1x12" transform="rotate(-72)"/><use width="180" height="120" href="#flagVe1x11" transform="rotate(144)"/></g></defs><path fill="#cf142b" d="M0 0h512v512H0z"/><path fill="#00247d" d="M0 0h512v341.3H0z"/><path fill="#fc0" d="M0 0h512v170.7H0z"/><g id="flagVe1x14" transform="translate(256.3 358.4) scale(4.265)"><g id="flagVe1x15"><use width="180" height="120" href="#flagVe1x10" transform="rotate(10)"/><use width="180" height="120" href="#flagVe1x10" transform="rotate(30)"/></g><use width="180" height="120" href="#flagVe1x15" transform="rotate(40)"/></g><use width="180" height="120" href="#flagVe1x14" transform="rotate(-80 256.3 358.4)"/>`),
		g.Group(children),
	)
}