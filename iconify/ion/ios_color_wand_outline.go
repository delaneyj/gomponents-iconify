package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosColorWandOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M192.011 149.661l-34.043 34.041 256.097 256.096L448 405.757 192.011 149.661zm20.585 66.041l11.415-11.414 201.468 201.469-11.414 11.414-201.469-201.469z" fill="currentColor"/><path d="M184 64h16v40h-16z" fill="currentColor"/><path d="M184 268h16v40h-16z" fill="currentColor"/><path d="M280 176h40v16h-40z" fill="currentColor"/><path d="M64 176h40v16H64z" fill="currentColor"/><path d="M139.675 122.558l-11.313 11.314-28.284-28.284 11.313-11.314z" fill="currentColor"/><path d="M111.389 278.128l-11.312-11.312 28.284-28.284 11.312 11.312z" fill="currentColor"/><path d="M255.641 133.877l-11.313-11.313L272.61 94.28l11.313 11.313z" fill="currentColor"/>`),
		g.Group(children),
	)
}