package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationMan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M49.798 23.592c-7.834.001-15.596 3.368-14.78 10.096l2 14.624c.351 2.573 2.09 6.688 4.687 6.688h.185l2.127 24.531c.092 1.104.892 2 2 2h8c1.108 0 1.908-.896 2-2L58.144 55h.186c2.597 0 4.335-4.115 4.687-6.688l2-14.624c.524-6.734-7.384-10.098-15.219-10.096Z" color="currentColor"/><path fill="currentColor" d="m50.024 50.908l-.048.126c.016-.038.027-.077.043-.115l.005-.011z" color="currentColor"/><circle cx="50" cy="10.5" r="10.5" fill="currentColor" color="currentColor"/><path fill="currentColor" d="M60.922 69.092c-.085.972-.175 1.942-.26 2.914C69.614 73.27 76.25 76.138 77 79.686c1.117 5.276-16.142 7.65-27.26 7.539c-11.118-.112-28.059-2.263-26.578-7.54c.972-3.463 7.512-6.274 16.23-7.583c-.087-.975-.186-1.95-.27-2.924c-11.206 1.236-20.542 4.279-24.272 8.246H2.229L0 82.047h13.166c1.023 5.44 12.427 10.136 28.734 11.322L41.342 100h16.14l-.162-6.63c16.39-1.187 28.117-5.883 29.514-11.323H100l-1.91-4.623H85.469c-3.543-4.067-13.048-7.16-24.547-8.332z" color="currentColor"/>`),
		g.Group(children),
	)
}