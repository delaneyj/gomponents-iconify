package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#f93" d="M0 0h512v170.7H0z"/><path fill="#fff" d="M0 170.7h512v170.6H0z"/><path fill="#128807" d="M0 341.3h512V512H0z"/><g transform="translate(256 256) scale(3.41333)"><circle r="20" fill="#008"/><circle r="17.5" fill="#fff"/><circle r="3.5" fill="#008"/><g id="flagIn1x10"><g id="flagIn1x11"><g id="flagIn1x12"><g id="flagIn1x13" fill="#008"><circle r=".9" transform="rotate(7.5 -8.8 133.5)"/><path d="M0 17.5L.6 7L0 2l-.6 5L0 17.5z"/></g><use width="100%" height="100%" href="#flagIn1x13" transform="rotate(15)"/></g><use width="100%" height="100%" href="#flagIn1x12" transform="rotate(30)"/></g><use width="100%" height="100%" href="#flagIn1x11" transform="rotate(60)"/></g><use width="100%" height="100%" href="#flagIn1x10" transform="rotate(120)"/><use width="100%" height="100%" href="#flagIn1x10" transform="rotate(-120)"/></g>`),
		g.Group(children),
	)
}