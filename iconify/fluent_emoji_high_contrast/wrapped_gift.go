package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrappedGift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.534 3.084C22.953.74 27 2.463 27 5.847c0 .424-.123.82-.336 1.153h.997A2.34 2.34 0 0 1 30 9.342v4.316a2.342 2.342 0 0 1-2 2.318v11.938C28 29.62 26.603 31 24.888 31H7.112C5.4 31 4 29.634 4 27.914V15.976a2.342 2.342 0 0 1-2-2.318V9.342A2.34 2.34 0 0 1 4.338 7h.993A2.154 2.154 0 0 1 5 5.847C5 2.463 9.045.731 11.466 3.084l1.79 1.736A3.003 3.003 0 0 1 16.006 3c1.225 0 2.278.745 2.742 1.816l1.787-1.732ZM6 16v11.914C6 28.507 6.482 29 7.112 29H13V16H6Zm13 0v13h5.888A1.1 1.1 0 0 0 26 27.914V16h-7Zm9-2.342V9.342A.34.34 0 0 0 27.662 9H19v5h8.662a.34.34 0 0 0 .338-.342ZM13 9H4.338A.34.34 0 0 0 4 9.342v4.316a.34.34 0 0 0 .338.342H13V9Z"/>`),
		g.Group(children),
	)
}