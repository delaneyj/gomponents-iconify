package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardTextBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.54 2.088C15.33 2 13.845 2 12 2h-.099l-3.5 5.25H13.1l3.441-5.162Zm-6.444-.084c-3.474.027-5.38.208-6.632 1.46c-.857.858-1.213 2.022-1.36 3.786H6.6l3.497-5.246Z"/><path fill-rule="evenodd" d="M2.026 8.75C2 9.689 2 10.763 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12c0-1.237 0-2.311-.026-3.25H2.026ZM5.75 14a.75.75 0 0 1 .75-.75h8a.75.75 0 0 1 0 1.5h-8a.75.75 0 0 1-.75-.75Zm.75 2.75a.75.75 0 0 0 0 1.5H12a.75.75 0 0 0 0-1.5H6.5Z" clip-rule="evenodd"/><path d="M20.535 3.464c.858.858 1.214 2.022 1.36 3.786h-6.994L18.2 2.304c.993.21 1.74.563 2.336 1.16Z"/></g>`),
		g.Group(children),
	)
}