package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timecap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.248 26.974L9.872 13.656c-2.71-3.85.043-9.156 4.752-9.156h18.752c4.709 0 7.462 5.306 4.752 9.156l-9.376 13.318c-2.315 3.288-7.189 3.288-9.504 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.248 21.026L9.872 34.344c-2.71 3.85.043 9.156 4.752 9.156h18.752c4.709 0 7.462-5.306 4.752-9.156l-9.376-13.318c-2.315-3.288-7.189-3.288-9.504 0Z"/>`),
		g.Group(children),
	)
}