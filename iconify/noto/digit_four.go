package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="ssvg-id-digit-foura" gradientUnits="userSpaceOnUse" x1="62.009" y1="102.808" x2="62.009" y2="27.26" gradientTransform="matrix(1 0 0 -1 0 128)"><stop offset="0" stop-color="#757575"/><stop offset=".515" stop-color="#504f4f"/></linearGradient><path d="M82.26 74.91h5.34c1.14 0 2.06.92 2.06 2.06v6.07c0 1.14-.92 2.06-2.06 2.06h-5.34c-1.14 0-2.06.92-2.06 2.06v12.8c0 1.14-.92 2.06-2.06 2.06h-8.57c-1.14 0-2.06-.92-2.06-2.06v-12.8c0-1.14-.92-2.06-2.06-2.06h-28.8c-1.1 0-2-.86-2.05-1.96l-.24-5.11c-.02-.42.09-.84.32-1.2L66.5 26.97a2.06 2.06 0 0 1 1.73-.95h9.92c1.14 0 2.06.92 2.06 2.06v44.8c-.01 1.11.91 2.03 2.05 2.03zm-30.9 0h14.11c1.14 0 2.06-.92 2.06-2.06V43.17l-.94 1.67l-16.97 26.92c-.87 1.37.12 3.15 1.74 3.15z" fill="url(#ssvg-id-digit-foura)"/>`),
		g.Group(children),
	)
}