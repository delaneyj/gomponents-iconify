package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPointFavouriteBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2c-4.418 0-8 3.646-8 8.143c0 4.462 2.553 9.67 6.537 11.531a3.45 3.45 0 0 0 2.926 0C17.447 19.812 20 14.605 20 10.144C20 5.645 16.418 2 12 2ZM9 8.757c0 1.02 1.165 2.097 2.043 2.764c.42.32.63.479.957.479c.328 0 .537-.16.957-.479C13.835 10.854 15 9.777 15 8.758c0-1.733-1.65-2.38-3-1.041c-1.35-1.339-3-.692-3 1.041Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}