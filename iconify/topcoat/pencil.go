package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M19.57 29.24L23.33 33s15.8-15.769 15.79-15.779c1.851-1.859 1.83-2.68 0-4.51c.01-.02-1.511-1.51-1.511-1.51L19.57 29.24zm-1.71-1.711L35.9 9.491l-3.551-3.55L14.31 23.98l3.55 3.549zM30.85 4.441l-1.5-1.5c-1.91-1.91-2.59-1.931-4.51 0L9.05 18.721l3.76 3.76l18.04-18.04zm-23.3 15.78L.5 41.5l21.33-7L7.55 20.221zm-1.5 10.638l5.26 5.262l-7.61 2.26l2.35-7.522z"/>`),
		g.Group(children),
	)
}