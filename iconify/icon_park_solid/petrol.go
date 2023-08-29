package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Petrol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPetrol0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 42V6h-3l-9 10H12l-6 6v20h36Z"/><path stroke="#fff" stroke-linecap="round" d="M12 16L22 6h18"/><path fill="#000" stroke="#000" stroke-linecap="round" d="M20.643 28.889c1.431-1.88 2.535-4.479 3.131-5.889c1.044 1.41 3.31 4.948 4.026 6.829c.894 2.35-1.342 5.171-4.026 5.171c-2.684 0-4.92-3.76-3.131-6.111Z"/><path stroke="#fff" stroke-linecap="round" d="m11 8l-7 7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPetrol0)"/>`),
		g.Group(children),
	)
}