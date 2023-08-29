package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.188 13.379l-6.011 6.01a5 5 0 0 1-7.072-7.07l8.486-8.486a3 3 0 1 1 4.243 4.242l-7.778 7.779a1.024 1.024 0 0 1-1.414 0a1.001 1.001 0 0 1 0-1.415l5.302-5.303a.5.5 0 1 0-.707-.707l-5.302 5.303l-.001.001a2 2 0 0 0 0 2.828a2.048 2.048 0 0 0 2.829 0l7.778-7.779a4 4 0 0 0-5.657-5.656l-8.486 8.485a6 6 0 0 0 4.244 10.243a5.957 5.957 0 0 0 4.242-1.758l6.01-6.01a.5.5 0 1 0-.707-.707z"/>`),
		g.Group(children),
	)
}