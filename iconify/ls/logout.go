package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 151v391c0 36 15 68 39 93c24 24 55 37 91 37h196v-81H130c-27 0-48-22-48-49V151c0-27 21-48 48-48h196V21H130c-36 0-67 14-91 38c-24 25-39 56-39 92zm215 118v156c0 18 16 33 34 33h181v123c0 11 6 20 16 25c4 1 8 1 10 1c7 0 13-2 18-7l235-235c11-9 10-27 0-37L474 94c-14-15-44-6-44 18v124H249c-18 0-34 15-34 33z"/>`),
		g.Group(children),
	)
}