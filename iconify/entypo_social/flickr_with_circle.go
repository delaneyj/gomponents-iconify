package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4C4.698.4.4 4.698.4 10s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6S15.302.4 10 .4zM7.436 12a1.99 1.99 0 0 1-1.982-2c0-1.105.887-2 1.982-2c1.094 0 1.982.895 1.982 2s-.889 2-1.982 2zm5.129 0a1.991 1.991 0 0 1-1.983-2c0-1.105.888-2 1.983-2a1.99 1.99 0 0 1 1.982 2c0 1.105-.887 2-1.982 2z"/>`),
		g.Group(children),
	)
}