package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Odometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896a384 384 0 1 0 0-768a384 384 0 0 0 0 768zm0 64a448 448 0 1 1 0-896a448 448 0 0 1 0 896z"/><path fill="currentColor" d="M192 512a320 320 0 1 1 640 0a32 32 0 1 1-64 0a256 256 0 1 0-512 0a32 32 0 0 1-64 0z"/><path fill="currentColor" d="M570.432 627.84A96 96 0 1 1 509.568 608l60.992-187.776A32 32 0 1 1 631.424 440l-60.992 187.776zM502.08 734.464a32 32 0 1 0 19.84-60.928a32 32 0 0 0-19.84 60.928z"/>`),
		g.Group(children),
	)
}