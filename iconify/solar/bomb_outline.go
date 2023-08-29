package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.717 2.078c-.436-1.104-1.998-1.104-2.434 0L16.66 3.66l-1.582.623c-1.104.436-1.104 1.998 0 2.434l.82.323l-1.119 1.12a8.25 8.25 0 1 0 1.06 1.06l1.12-1.119l.324.821c.436 1.104 1.998 1.104 2.434 0l.623-1.582l1.582-.623c1.104-.436 1.104-1.998 0-2.434L20.34 3.66l-.623-1.582Zm-1.693 2.21l.476-1.206l.476 1.206c.133.337.4.603.736.736l1.206.476l-1.206.476c-.337.133-.603.4-.736.736L18.5 7.918l-.476-1.206a1.308 1.308 0 0 0-.736-.736L16.082 5.5l1.206-.476c.337-.133.603-.4.736-.736ZM9.5 7.75a6.75 6.75 0 1 0 0 13.5a6.75 6.75 0 0 0 0-13.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}