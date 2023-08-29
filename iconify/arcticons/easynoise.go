package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easynoise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.642 42.5l-13.82-11.834H5.5V17.334h9.323L28.643 5.5Zm8.493-31.33s5.365.813 5.364 2.388c-.002 2.303-7.85 1.178-7.85 3.48s7.85 1.179 7.85 3.482S34.65 21.698 34.65 24s7.849 1.179 7.849 3.482s-7.758 1.18-7.849 3.48c-.094 2.39 7.966 1.59 7.849 3.98c-.073 1.483-5.102 2.187-5.102 2.187"/>`),
		g.Group(children),
	)
}