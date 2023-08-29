package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFiveFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 43q18 0 30.5 12.5T384 85v214q0 17-12.5 29.5T341 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h298zm0 256V85H43v214h298z"/>`),
		g.Group(children),
	)
}