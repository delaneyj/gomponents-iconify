package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GdOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagGd1x10"><g id="flagGd1x11"><path id="flagGd1x12" fill="#fcd116" d="M0-1v1h.5" transform="rotate(18 0 -1)"/><use width="100%" height="100%" href="#flagGd1x12" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagGd1x11" transform="rotate(72)"/><use width="100%" height="100%" href="#flagGd1x11" transform="rotate(144)"/><use width="100%" height="100%" href="#flagGd1x11" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagGd1x11" transform="rotate(-72)"/></g></defs><path fill="#ce1126" d="M0 0h512v512H0z"/><path fill="#007a5e" d="M71.7 71.7h368.6v368.6H71.7z"/><path fill="#fcd116" d="M71.7 71.7h368.6L71.7 440.4h368.6z"/><circle cx="255.9" cy="256.1" r="61.4" fill="#ce1126"/><use width="100%" height="100%" href="#flagGd1x10" transform="translate(256 256) scale(56.32)"/><use width="100%" height="100%" x="-100" href="#flagGd1x13" transform="translate(-16.4 -.1)"/><use id="flagGd1x13" width="100%" height="100%" href="#flagGd1x10" transform="translate(256 35.9) scale(33.28)"/><use width="100%" height="100%" x="100" href="#flagGd1x13" transform="translate(16.4)"/><path fill="#ce1126" d="M99.8 256.8c7.7 14.3 22.6 29.8 35.7 35.3c.2-14.5-5-33.2-12-48l-23.7 12.7z"/><path fill="#fcd116" d="M86.8 207.6c11.1 23.3-29 78.7 37.8 91.7a67.5 67.5 0 0 1-11.5-44.7a75.5 75.5 0 0 1 34.6 32.8c17.5-63.4-44.8-59.5-61-79.8z"/><use width="100%" height="100%" x="-100" href="#flagGd1x13" transform="translate(-16.4 442)"/><use width="100%" height="100%" href="#flagGd1x10" transform="translate(256 478) scale(33.28)"/><use width="100%" height="100%" x="100" href="#flagGd1x13" transform="translate(16.4 442.2)"/>`),
		g.Group(children),
	)
}