package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 36V12a3 3 0 0 1 5.305-1.92L31 27.713V12a3 3 0 0 1 6 0v24a3 3 0 0 1-5.305 1.92L17 20.286V36a3 3 0 1 1-6 0Zm2-24v24a1 1 0 1 0 2 0V17.524a1 1 0 0 1 1.768-.64L33.232 36.64A1 1 0 0 0 35 36V12a1 1 0 0 0-2 0v18.476a1 1 0 0 1-1.768.64L14.768 11.36l.768-.64l-.768.64A1 1 0 0 0 13 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}