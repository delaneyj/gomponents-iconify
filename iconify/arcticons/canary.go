package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.472 16.938l25.712 7.595s-1.75 3.425-6.837 3.99"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.926 38.522S12.278 42.019 6.65 30.25s.833-21.916 7.44-21.91c5.881.007 9.725 5.478 9.725 5.478s-9.715.987-8.143 10.622s12.834 17.62 22.853 14.553L17.472 16.938L43.5 32.913a11.31 11.31 0 0 1-8.303 2.593"/>`),
		g.Group(children),
	)
}