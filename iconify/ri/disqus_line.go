package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisqusLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.997 2c5.523 0 10 4.477 10 10s-4.477 10-10 10a9.961 9.961 0 0 1-6.249-2.192l-4.718.59l1.72-4.586A9.972 9.972 0 0 1 1.996 12c0-5.523 4.477-10 10-10Zm0 2a8 8 0 0 0-7.262 11.362l.177.38l-.848 2.26l2.315-.289l.338.297a7.965 7.965 0 0 0 5.28 1.99a8 8 0 1 0 0-16Zm-3.95 3h3.79c3.42 0 5.44 1.956 5.54 4.729l.004.215v.027c0 2.814-1.963 4.922-5.338 5.025l-.262.004H8.047V7h3.79h-3.79Zm3.832 2.458H10.77v5.085h1.109c1.565 0 2.624-.845 2.703-2.345l.005-.183v-.028c0-1.6-1.08-2.53-2.708-2.53Z"/>`),
		g.Group(children),
	)
}