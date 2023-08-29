package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#f93" d="M0 0h640v160H0z"/><path fill="#fff" d="M0 160h640v160H0z"/><path fill="#128807" d="M0 320h640v160H0z"/><g transform="matrix(3.2 0 0 3.2 320 240)"><circle r="20" fill="#008"/><circle r="17.5" fill="#fff"/><circle r="3.5" fill="#008"/><g id="flagIn4x30"><g id="flagIn4x31"><g id="flagIn4x32"><g id="flagIn4x33" fill="#008"><circle r=".9" transform="rotate(7.5 -8.8 133.5)"/><path d="M0 17.5L.6 7L0 2l-.6 5L0 17.5z"/></g><use width="100%" height="100%" href="#flagIn4x33" transform="rotate(15)"/></g><use width="100%" height="100%" href="#flagIn4x32" transform="rotate(30)"/></g><use width="100%" height="100%" href="#flagIn4x31" transform="rotate(60)"/></g><use width="100%" height="100%" href="#flagIn4x30" transform="rotate(120)"/><use width="100%" height="100%" href="#flagIn4x30" transform="rotate(-120)"/></g>`),
		g.Group(children),
	)
}