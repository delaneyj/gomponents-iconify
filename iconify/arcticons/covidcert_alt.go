package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidcertAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.15 5.8c2.816-1.412 9.818-1.443 12.628 0c-.002 4.445-.005 10.364-6.314 13.26c-6.3-2.95-6.314-8.818-6.314-13.26Zm24.246 34.728H13.99a2.841 2.841 0 0 1-2.526-2.526V19.06m6.314-11.365h16.416a2.841 2.841 0 0 1 2.526 2.526v17.047"/><circle cx="34.573" cy="34.845" r="7.577" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.199 35.23l2.75 2.755l5.964-6.185m-8.507-6.426h-5.051v-5.05h5.051ZM17.778 30.3h10.734m-8.84 4.166h7.135m-3.977-21.72v5.051h-5.052m7.577-5.051h5.051v5.051h-5.051Zm-7.577 7.577h5.051v5.051h-5.051ZM11.464 7.442v7.577m-3.788-3.914h7.576"/>`),
		g.Group(children),
	)
}