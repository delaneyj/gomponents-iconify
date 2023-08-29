package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M421.5 602.5Q384 565 384 512t38-91L966 59L603 603q-37 37-90.5 37t-91-37.5zM512 62q-91 0-174.5 36T194 194T98.5 337T63 511q0 106 48 202l-56 28Q0 632 0 512q0-105 40.5-199.5t109-163T313 40.5T512 0q161 0 293 92l-57 38Q639 62 512 62zm421 158q91 132 91 292q0 120-55 229l-56-29q48-95 48-201q0-125-66-234z"/>`),
		g.Group(children),
	)
}