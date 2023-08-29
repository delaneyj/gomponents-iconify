package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#635994" d="M7.223 2.893a3.072 3.072 0 0 0-4.33 0a3.084 3.084 0 0 0 0 4.34l8.747 8.744l-8.738 8.745a3.072 3.072 0 0 0 0 4.33a3.072 3.072 0 0 0 4.33 0l8.741-8.744l8.74 8.735a3.072 3.072 0 0 0 4.33 0a3.072 3.072 0 0 0 0-4.33l-8.738-8.739l8.747-8.752a3.072 3.072 0 0 0 0-4.33c-1.2-1.19-3.15-1.19-4.34 0l-8.74 8.75l-8.75-8.75Z"/>`),
		g.Group(children),
	)
}