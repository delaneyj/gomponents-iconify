package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 138h-80a6 6 0 0 0-6 6v57.45a6 6 0 0 0 4.93 5.91l80 14.54a5.46 5.46 0 0 0 1.07.1a6 6 0 0 0 6-6v-72a6 6 0 0 0-6-6Zm-6 70.81l-68-12.36V150h68ZM104 138H40a6 6 0 0 0-6 6v40a6 6 0 0 0 4.93 5.9l64 11.64a6.36 6.36 0 0 0 1.07.1a6 6 0 0 0 6-6V144a6 6 0 0 0-6-6Zm-6 50.45L46 179v-29h52ZM219.84 35.39a6 6 0 0 0-4.91-1.29l-80 14.54a6 6 0 0 0-4.93 5.91V112a6 6 0 0 0 6 6h80a6 6 0 0 0 6-6V40a6 6 0 0 0-2.16-4.61ZM210 106h-68V59.55l68-12.36ZM107.84 55.76a6 6 0 0 0-4.91-1.3l-64 11.64A6 6 0 0 0 34 72v40a6 6 0 0 0 6 6h64a6 6 0 0 0 6-6V60.36a6 6 0 0 0-2.16-4.6ZM98 106H46V77l52-9.46Z"/>`),
		g.Group(children),
	)
}