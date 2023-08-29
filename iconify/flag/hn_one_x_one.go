package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#18c3df" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 170.7h512v170.6H0z"/><g id="flagHn1x10" fill="#18c3df" transform="translate(256 256) scale(28.44446)"><g id="flagHn1x11"><path id="flagHn1x12" d="m0-1l-.3 1l.5.1z"/><use width="100%" height="100%" href="#flagHn1x12" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagHn1x11" transform="rotate(72)"/><use width="100%" height="100%" href="#flagHn1x11" transform="rotate(-72)"/><use width="100%" height="100%" href="#flagHn1x11" transform="rotate(144)"/><use width="100%" height="100%" href="#flagHn1x11" transform="rotate(-144)"/></g><use width="100%" height="100%" href="#flagHn1x10" transform="translate(142.2 -45.5)"/><use width="100%" height="100%" href="#flagHn1x10" transform="translate(142.2 39.8)"/><use width="100%" height="100%" href="#flagHn1x10" transform="translate(-142.2 -45.5)"/><use width="100%" height="100%" href="#flagHn1x10" transform="translate(-142.2 39.8)"/>`),
		g.Group(children),
	)
}