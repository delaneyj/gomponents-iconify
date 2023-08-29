package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func University(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256 64L32 176.295l32 21.072V400l32 16V217.849L255.852 320l132.144-85.451L480 174.727z" fill="currentColor"/><path d="M390.13 256L256 343.768 120.531 256 112 337.529C128 349.984 224 416 256.002 448 288 416 384 350.031 400 337.561L390.13 256z" fill="currentColor"/>`),
		g.Group(children),
	)
}