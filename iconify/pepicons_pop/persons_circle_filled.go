package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonsCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPersonsCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.236 13.055a4.954 4.954 0 0 0-.832 2.748v.697a1 1 0 0 1-2 0v-.697c0-1.373.406-2.715 1.168-3.858l1.664 1.11ZM9.207 12c-.792 0-1.532.396-1.971 1.055l-1.664-1.11A4.369 4.369 0 0 1 9.207 10h.197a1 1 0 0 1 0 2h-.197Zm2.211 1.055c.543.813.832 1.77.832 2.748v.697a1 1 0 0 0 2 0v-.697a6.955 6.955 0 0 0-1.168-3.858l-1.664 1.11Z"/><path d="M9.447 12c.792 0 1.532.396 1.971 1.055l1.664-1.11A4.369 4.369 0 0 0 9.447 10H9.25a1 1 0 0 0 0 2h.197Z"/><path d="M9.25 9.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm0 2a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5Zm5.236 5.055a4.954 4.954 0 0 0-.832 2.748v.697a1 1 0 0 1-2 0v-.697c0-1.373.406-2.715 1.168-3.858l1.664 1.11Zm1.971-1.055c-.792 0-1.532.396-1.971 1.055l-1.664-1.11a4.369 4.369 0 0 1 3.635-1.945h.197a1 1 0 0 1 0 2h-.197Zm2.211 1.055c.543.813.832 1.77.832 2.748v.697a1 1 0 0 0 2 0v-.697a6.955 6.955 0 0 0-1.168-3.858l-1.664 1.11Z"/><path d="M16.697 15.25c.792 0 1.532.396 1.971 1.055l1.664-1.11a4.369 4.369 0 0 0-3.635-1.945H16.5a1 1 0 1 0 0 2h.197Z"/><path d="M16.5 12.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm0 2a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPersonsCircleFilled0)"/></g>`),
		g.Group(children),
	)
}