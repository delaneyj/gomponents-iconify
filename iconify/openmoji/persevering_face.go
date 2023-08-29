package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerseveringFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M36.2 13.316c-12.572 0-22.8 10.228-22.8 22.8c0 12.572 10.228 22.8 22.8 22.8c12.572 0 22.8-10.228 22.8-22.8c0-12.572-10.228-22.8-22.8-22.8z"/><path fill="#FFF" d="M28 51c.27-.356 3.31-8.218 8.421-8.004C41.026 43.19 43.65 50.537 44 51H28z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23"/><path stroke-linecap="round" stroke-linejoin="round" d="m44.536 21.439l6.385 3.277m-24-3.277l-6.385 3.277M24 28l7 4l-7 4m24-8l-7 4l7 4M28 51c.27-.356 1-8 8.421-8.004C43 42.993 43.65 50.537 44 51H28z"/></g>`),
		g.Group(children),
	)
}