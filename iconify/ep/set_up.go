package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M224 160a64 64 0 0 0-64 64v576a64 64 0 0 0 64 64h576a64 64 0 0 0 64-64V224a64 64 0 0 0-64-64H224zm0-64h576a128 128 0 0 1 128 128v576a128 128 0 0 1-128 128H224A128 128 0 0 1 96 800V224A128 128 0 0 1 224 96z"/><path fill="currentColor" d="M384 416a64 64 0 1 0 0-128a64 64 0 0 0 0 128zm0 64a128 128 0 1 1 0-256a128 128 0 0 1 0 256z"/><path fill="currentColor" d="M480 320h256q32 0 32 32t-32 32H480q-32 0-32-32t32-32zm160 416a64 64 0 1 0 0-128a64 64 0 0 0 0 128zm0 64a128 128 0 1 1 0-256a128 128 0 0 1 0 256z"/><path fill="currentColor" d="M288 640h256q32 0 32 32t-32 32H288q-32 0-32-32t32-32z"/>`),
		g.Group(children),
	)
}