package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blutspende(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.775 42.781C18.477 45.971 3.8 38.464 13.586 21.786c3.416-5.822 6.709-11.722 10.51-17.286C32.43 17.15 39.83 29.077 37.013 35.384"/>`),
		g.Group(children),
	)
}