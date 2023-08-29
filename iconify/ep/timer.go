package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896a320 320 0 1 0 0-640a320 320 0 0 0 0 640zm0 64a384 384 0 1 1 0-768a384 384 0 0 1 0 768z"/><path fill="currentColor" d="M512 320a32 32 0 0 1 32 32l-.512 224a32 32 0 1 1-64 0L480 352a32 32 0 0 1 32-32z"/><path fill="currentColor" d="M448 576a64 64 0 1 0 128 0a64 64 0 1 0-128 0zm96-448v128h-64V128h-96a32 32 0 0 1 0-64h256a32 32 0 1 1 0 64h-96z"/>`),
		g.Group(children),
	)
}