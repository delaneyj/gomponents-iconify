package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M38 13s2 18 2 23c0 4.123-2.172 7.586-5 9c-4 2-14 7-19 7h-3.4a4.894 4.894 0 0 0-.6 2.5a3.5 3.5 0 0 0 3.5 3.5H49a5 5 0 0 0 5-5c0-6-3-11-3-17s3-23 3-23v-1H38Z"/><path fill="#f1b31c" d="M45.5 58H49a5 5 0 0 0 5-5c0-6-3-11-3-17s3-23 3-23v-1h-4s-3 17-3 23s3 11 3 17a5 5 0 0 1-5 5Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M38 13s2 18 2 23c0 4.123-2.172 7.586-5 9c-4 2-14 7-19 7h-3.4a4.894 4.894 0 0 0-.6 2.5a3.5 3.5 0 0 0 3.5 3.5H49a5 5 0 0 0 5-5c0-6-3-11-3-17s3-23 3-23"/><path stroke-miterlimit="10" d="M48 47a4 4 0 0 1-4 4"/></g>`),
		g.Group(children),
	)
}