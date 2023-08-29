package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeLightroomClassic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 .401v31.197h32V.401zm1.333 1.334h29.333v28.531H1.333zm7.531 20.432c-.156 0-.197-.068-.197-.197V7.646c0-.115.041-.176.156-.176h2.109c.115 0 .136.047.136.161v12.343h5.5c.109 0 .14.048.12.157l-.333 1.891c-.021.109-.089.156-.199.156H8.864zm9.871-8.636c0-.161 0-.552-.068-1.301c0-.109.027-.136.115-.177c.812-.328 2.74-.921 4.896-.921c.109 0 .151.02.151.135v1.932c0 .115-.041.136-.156.136c-.833-.041-2.084.068-2.547.265v8.412c0 .109-.043.151-.152.151h-2.093c-.109 0-.151-.041-.151-.151v-8.48z"/>`),
		g.Group(children),
	)
}