package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookMarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.75 8A4.75 4.75 0 0 1 8.5 3.25h10c.966 0 1.75.784 1.75 1.75v15a1.75 1.75 0 0 1-1.75 1.75h-11A3.75 3.75 0 0 1 3.75 18V8Zm15-3v9.25H7.5c-.844 0-1.623.279-2.25.75V8A3.25 3.25 0 0 1 8.5 4.75h3.208a20.793 20.793 0 0 0 .071 5.97l.063.397a.75.75 0 0 0 1.292.392l1.366-1.48l1.366 1.48a.75.75 0 0 0 1.292-.392l.063-.397a20.79 20.79 0 0 0 .071-5.97H18.5a.25.25 0 0 1 .25.25Zm-2.972-.25h-2.556a19.293 19.293 0 0 0-.108 4.57l.651-.706a1 1 0 0 1 1.47 0l.651.706a19.293 19.293 0 0 0-.108-4.57Zm-8.278 11h11.25V20a.25.25 0 0 1-.25.25h-11a2.25 2.25 0 0 1 0-4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}