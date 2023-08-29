package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#fcc21b"><path d="M119.175 9.018v41.26h-16.42V9.018z"/><path d="M119.278 25.432h-41.26V9.012h41.26z"/><path d="M110.389 29.414L78.873 60.93l-11.61-11.61l31.515-31.516z"/><path d="M77.49 50.7c-15.96-15.96-41.84-15.96-57.8 0c-15.96 15.96-15.96 41.84 0 57.8s41.84 15.96 57.8 0c15.96-15.96 15.96-41.84 0-57.8zM31.61 96.58c-9.38-9.38-9.38-24.58 0-33.96c9.38-9.38 24.58-9.38 33.96 0s9.38 24.58 0 33.96c-9.37 9.38-24.58 9.38-33.96 0z"/></g>`),
		g.Group(children),
	)
}