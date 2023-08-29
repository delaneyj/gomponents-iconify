package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagEu4x30"><g id="flagEu4x31"><path id="flagEu4x32" d="m0-1l-.3 1l.5.1z"/><use href="#flagEu4x32" transform="scale(-1 1)"/></g><g id="flagEu4x33"><use href="#flagEu4x31" transform="rotate(72)"/><use href="#flagEu4x31" transform="rotate(144)"/></g><use href="#flagEu4x33" transform="scale(-1 1)"/></g></defs><path fill="#039" d="M0 0h640v480H0z"/><g fill="#fc0" transform="translate(320 242.3) scale(23.7037)"><use width="100%" height="100%" y="-6" href="#flagEu4x30"/><use width="100%" height="100%" y="6" href="#flagEu4x30"/><g id="flagEu4x34"><use width="100%" height="100%" x="-6" href="#flagEu4x30"/><use width="100%" height="100%" href="#flagEu4x30" transform="rotate(-144 -2.3 -2.1)"/><use width="100%" height="100%" href="#flagEu4x30" transform="rotate(144 -2.1 -2.3)"/><use width="100%" height="100%" href="#flagEu4x30" transform="rotate(72 -4.7 -2)"/><use width="100%" height="100%" href="#flagEu4x30" transform="rotate(72 -5 .5)"/></g><use width="100%" height="100%" href="#flagEu4x34" transform="scale(-1 1)"/></g>`),
		g.Group(children),
	)
}