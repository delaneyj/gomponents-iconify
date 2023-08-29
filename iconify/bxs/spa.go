package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.114c-3.998-5.951-8.574-7.043-8.78-7.09L2 8.75V10c0 7.29 3.925 12 10 12c5.981 0 10-4.822 10-12V8.75l-1.22.274c-.206.047-4.782 1.139-8.78 7.09z"/><path fill="currentColor" d="M11.274 3.767c-1.799 1.898-2.84 3.775-3.443 5.295c1.329.784 2.781 1.943 4.159 3.685c1.364-1.76 2.826-2.925 4.17-3.709c-.605-1.515-1.646-3.383-3.435-5.271L12 3l-.726.767z"/>`),
		g.Group(children),
	)
}