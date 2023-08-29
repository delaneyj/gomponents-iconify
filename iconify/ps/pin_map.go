package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 0Q101 0 53 53.5T5 192q0 95 154 311l17 24l17-24q154-216 154-311q0-85-48-138.5T176 0zm0 454q-47-68-87.5-144.5T48 192q0-61 33.5-105T176 43t94.5 44T304 192q0 41-40.5 117.5T176 454zm0-369q-35 0-60 25.5T91 171t25 60t60 25t60-25t25-60t-25-60.5T176 85zm0 128q-17 0-30-12.5T133 171q0-18 13-30.5t30-12.5t30 12.5t13 30.5q0 17-13 29.5T176 213z"/>`),
		g.Group(children),
	)
}