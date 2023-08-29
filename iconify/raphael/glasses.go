package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.075 9.53S11.37 8.094 8.917 8.094c-2.453 0-4.862.593-4.862.593L3.97 9.87l.53.53c.337.335.485 3.74 1.837 5.093c1.353 1.354 4.82 1.396 5.963.676c1.14-.72 2.24-3.467 2.24-4.694V10.8c.275-.275 1.616-.303 1.918 0v.676c0 1.227 1.1 3.975 2.24 4.693c1.145.72 4.612.677 5.964-.677c1.355-1.353 1.5-4.757 1.84-5.094l.527-.53l-.085-1.184s-2.408-.593-4.862-.593c-2.453 0-5.158 1.438-5.158 1.438c-.606-.238-2.188-.21-2.85 0z"/>`),
		g.Group(children),
	)
}