package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.095 2.004c-3.473.027-5.378.208-6.63 1.46c-.858.858-1.214 2.022-1.361 3.786h4.494l3.497-5.246ZM21.896 7.25c-.148-1.764-.503-2.928-1.36-3.786c-.598-.597-1.344-.95-2.338-1.16L14.901 7.25h6.995ZM16.54 2.088C15.33 2 13.845 2 12 2h-.1L8.4 7.25h4.697l3.442-5.162Z"/><path d="M2.026 8.75C2 9.689 2 10.763 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12c0-1.237 0-2.311-.026-3.25H2.026Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}