package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagEu1x10"><g id="flagEu1x11"><path id="flagEu1x12" d="m0-1l-.3 1l.5.1z"/><use href="#flagEu1x12" transform="scale(-1 1)"/></g><g id="flagEu1x13"><use href="#flagEu1x11" transform="rotate(72)"/><use href="#flagEu1x11" transform="rotate(144)"/></g><use href="#flagEu1x13" transform="scale(-1 1)"/></g></defs><path fill="#039" d="M0 0h512v512H0z"/><g fill="#fc0" transform="translate(256 258.4) scale(25.28395)"><use width="100%" height="100%" y="-6" href="#flagEu1x10"/><use width="100%" height="100%" y="6" href="#flagEu1x10"/><g id="flagEu1x14"><use width="100%" height="100%" x="-6" href="#flagEu1x10"/><use width="100%" height="100%" href="#flagEu1x10" transform="rotate(-144 -2.3 -2.1)"/><use width="100%" height="100%" href="#flagEu1x10" transform="rotate(144 -2.1 -2.3)"/><use width="100%" height="100%" href="#flagEu1x10" transform="rotate(72 -4.7 -2)"/><use width="100%" height="100%" href="#flagEu1x10" transform="rotate(72 -5 .5)"/></g><use width="100%" height="100%" href="#flagEu1x14" transform="scale(-1 1)"/></g>`),
		g.Group(children),
	)
}