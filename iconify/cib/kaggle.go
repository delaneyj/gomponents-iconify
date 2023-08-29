package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.099 31.812c-.025.12-.156.188-.375.188h-4.183c-.249 0-.468-.109-.656-.328l-6.907-8.787l-1.932 1.828v6.817c0 .313-.151.469-.463.469H7.338c-.312 0-.469-.156-.469-.469V.469c0-.308.157-.469.469-.469h3.245c.312 0 .463.161.463.469v19.124l8.271-8.359c.224-.224.443-.328.661-.328h4.319c.192 0 .317.077.38.239c.063.199.047.339-.047.417l-8.74 8.459l9.115 11.343c.125.141.156.276.093.48z"/>`),
		g.Group(children),
	)
}