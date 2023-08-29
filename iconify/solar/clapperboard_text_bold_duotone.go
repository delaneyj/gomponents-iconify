package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardTextBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.54 2.088C15.33 2 13.846 2 12 2h-.098l-3.5 5.25h4.697l3.441-5.162Zm-6.444-.084c-3.474.027-5.379.208-6.631 1.46c-.858.858-1.213 2.022-1.36 3.786h4.494l3.497-5.246Zm10.44 1.46c.857.858 1.213 2.022 1.36 3.786h-6.994l3.297-4.946c.994.21 1.74.563 2.337 1.16Z"/><path d="M2 12c0-1.237 0-2.311.026-3.25h19.948C22 9.689 22 10.763 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z" opacity=".5"/><path d="M5.75 17.5a.75.75 0 0 1 .75-.75H12a.75.75 0 0 1 0 1.5H6.5a.75.75 0 0 1-.75-.75Zm.75-4.25a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5h-8Z"/></g>`),
		g.Group(children),
	)
}