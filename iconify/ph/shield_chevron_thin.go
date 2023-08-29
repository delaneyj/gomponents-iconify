package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldChevronThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 44H48a12 12 0 0 0-12 12v58.77c0 86.87 73.54 115.7 88.28 120.6a11.65 11.65 0 0 0 7.44 0c14.74-4.9 88.28-33.73 88.28-120.6V56a12 12 0 0 0-12-12Zm-78.8 183.79a3.53 3.53 0 0 1-2.4 0c-9.75-3.24-47-17.66-68-54.45l69.2-48.46l69.23 48.46c-21.07 36.79-58.23 51.21-68.03 54.45Zm82.8-113c0 20.38-4.33 37.35-11 51.41l-70.68-49.48a4 4 0 0 0-4.58 0L55 166.2c-6.7-14.06-11-31-11-51.41V56a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}