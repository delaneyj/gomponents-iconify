package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HnFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#18c3df" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 160h640v160H0z"/><g id="flagHn4x30" fill="#18c3df" transform="translate(320 240) scale(26.66665)"><g id="flagHn4x31"><path id="flagHn4x32" d="m-.3 0l.5.1L0-1z"/><use width="100%" height="100%" href="#flagHn4x32" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagHn4x31" transform="rotate(72)"/><use width="100%" height="100%" href="#flagHn4x31" transform="rotate(-72)"/><use width="100%" height="100%" href="#flagHn4x31" transform="rotate(144)"/><use width="100%" height="100%" href="#flagHn4x31" transform="rotate(-144)"/></g><use width="100%" height="100%" href="#flagHn4x30" transform="translate(133.3 -42.7)"/><use width="100%" height="100%" href="#flagHn4x30" transform="translate(133.3 37.3)"/><use width="100%" height="100%" href="#flagHn4x30" transform="translate(-133.3 -42.7)"/><use width="100%" height="100%" href="#flagHn4x30" transform="translate(-133.3 37.3)"/>`),
		g.Group(children),
	)
}