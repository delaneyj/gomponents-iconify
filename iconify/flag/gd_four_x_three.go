package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GdFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagGd4x30"><g id="flagGd4x31"><path id="flagGd4x32" fill="#fcd116" d="M0-1v1h.5" transform="rotate(18 0 -1)"/><use href="#flagGd4x32" transform="scale(-1 1)"/></g><use href="#flagGd4x31" transform="rotate(72)"/><use href="#flagGd4x31" transform="rotate(144)"/><use href="#flagGd4x31" transform="rotate(216)"/><use href="#flagGd4x31" transform="rotate(288)"/></g></defs><path fill="#ce1126" d="M0 0h640v480H0z"/><path fill="#007a5e" d="M67.2 67.2h505.6v345.6H67.2z"/><path fill="#fcd116" d="M67.2 67.3h505.6L67.2 412.9h505.6z"/><circle cx="319.9" cy="240.1" r="57.6" fill="#ce1126"/><use width="100%" height="100%" href="#flagGd4x30" transform="matrix(52.8 0 0 52.8 320 240)"/><use width="100%" height="100%" x="-100" href="#flagGd4x33" transform="translate(-30.3)"/><use id="flagGd4x33" width="100%" height="100%" href="#flagGd4x30" transform="matrix(31.2 0 0 31.2 320 33.6)"/><use width="100%" height="100%" x="100" href="#flagGd4x33" transform="translate(30.3)"/><path fill="#ce1126" d="M102.3 240.7a80.4 80.4 0 0 0 33.5 33.2a111 111 0 0 0-11.3-45l-22.2 11.8z"/><path fill="#fcd116" d="M90.1 194.7c10.4 21.7-27.1 73.7 35.5 85.9a63.2 63.2 0 0 1-10.9-41.9a70 70 0 0 1 32.5 30.8c16.4-59.5-42-55.8-57.1-74.8z"/><use width="100%" height="100%" x="-100" href="#flagGd4x33" transform="translate(-30.3 414.6)"/><use width="100%" height="100%" href="#flagGd4x30" transform="matrix(31.2 0 0 31.2 320 448.2)"/><use width="100%" height="100%" x="100" href="#flagGd4x33" transform="translate(30.3 414.6)"/>`),
		g.Group(children),
	)
}