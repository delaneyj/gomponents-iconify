package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#377e3f" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 102.4h512v307.2H0z"/><path fill="#b40a2d" d="M0 153.6h512v204.8H0z"/><path fill="#ecc81d" d="m255.9 163.4l60.2 185.2l-157.6-114.5h194.8L195.7 348.6z"/>`),
		g.Group(children),
	)
}