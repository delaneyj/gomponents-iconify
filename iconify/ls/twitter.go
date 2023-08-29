package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M768 108c-22 32-48 59-79 82c1 7 1 14 1 21c0 208-158 448-448 448c-89 0-172-27-242-71c12 1 25 3 38 3c74 0 141-26 195-68c-70-1-127-48-147-110c9 2 20 3 30 3c14 0 29-2 42-6c-73-14-127-77-127-153v-2c22 11 45 18 71 19c-42-29-70-77-70-131c0-29 8-56 21-80c78 95 194 159 325 165c-2-11-4-24-4-36c0-87 70-158 157-158c46 0 87 20 116 51c36-7 69-21 99-39c-11 37-36 69-68 88c32-4 61-13 90-26z"/>`),
		g.Group(children),
	)
}