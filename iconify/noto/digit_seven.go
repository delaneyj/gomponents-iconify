package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-digit-sevena" gradientUnits="userSpaceOnUse" x1="64" y1="102.5" x2="64" y2="24.97" gradientTransform="matrix(1 0 0 -1 0 128)"><stop offset="0" stop-color="#757575"/><stop offset=".515" stop-color="#504f4f"/></linearGradient><path d="M90.16 33.44L60.5 100.77c-.33.75-1.07 1.23-1.88 1.23h-8.76c-1.5 0-2.49-1.55-1.87-2.92l27.64-59.99c.63-1.36-.37-2.92-1.87-2.92H39.72c-1.14 0-2.06-.92-2.06-2.06v-6.07c0-1.14.92-2.06 2.06-2.06h48.56c1.14 0 2.06.92 2.06 2.06v4.56c-.01.3-.07.58-.18.84z" fill="url(#ssvg-id-digit-sevena)"/>`),
		g.Group(children),
	)
}