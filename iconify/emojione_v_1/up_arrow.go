package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.667 57.064c0 3.812-3.07 6.9-6.86 6.9H6.86c-3.788 0-6.86-3.088-6.86-6.9V6.896C0 3.088 3.072 0 6.86 0h49.946c3.79 0 6.86 3.088 6.86 6.896v50.168z"/><path fill="#fff" d="M15.937 28.09L31.795 9.578l15.938 18.334c-.191 2.574-1.894 4.593-3.981 4.603l-3.54.018l.09 16.704c.012 2.81-2.245 5.1-5.04 5.117l-6.574.034c-2.798.012-5.079-2.256-5.092-5.065l-.088-16.7l-3.538.02c-2.086.007-3.808-1.985-4.03-4.553"/>`),
		g.Group(children),
	)
}