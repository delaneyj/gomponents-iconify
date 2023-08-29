package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DishwasherThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5V12H3V7.5Zm8.768 6.732A2.513 2.513 0 0 0 11.5 14H29v10.5a4.5 4.5 0 0 1-4.5 4.5h-14A2.5 2.5 0 0 0 9 26.708v-2.002a6.001 6.001 0 0 0 3.5-5.456V16a2.5 2.5 0 0 0-.732-1.768ZM11.5 8a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0ZM17 7a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2h-5ZM3 15a1 1 0 0 0-1 1v3.25a4.502 4.502 0 0 0 3.5 4.389V28H5a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2h-.5v-4.361A4.502 4.502 0 0 0 11 19.25V16a1 1 0 0 0-1-1H3Z"/>`),
		g.Group(children),
	)
}