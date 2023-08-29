package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardPlayBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 12c0-1.237 0-2.311.026-3.25h19.948C22 9.689 22 10.763 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z" clip-rule="evenodd" opacity=".5"/><path d="M15 14.5c0-.633-.662-1.06-1.986-1.915c-1.342-.866-2.013-1.299-2.514-.98c-.5.317-.5 1.176-.5 2.895s0 2.578.5 2.896c.5.318 1.172-.115 2.514-.981C14.338 15.56 15 15.133 15 14.5ZM12 2c1.845 0 3.33 0 4.54.088L13.098 7.25H8.401l3.5-5.25H12ZM3.464 3.464c1.253-1.252 3.158-1.433 6.631-1.46L6.598 7.25H2.104c.147-1.764.503-2.928 1.36-3.786ZM21.896 7.25c-.148-1.764-.503-2.928-1.36-3.786c-.598-.597-1.344-.95-2.338-1.16L14.901 7.25h6.995Z"/></g>`),
		g.Group(children),
	)
}