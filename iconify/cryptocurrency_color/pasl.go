package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pasl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#00acff"/><path fill="#fff" d="M15.08 20.156a.51.51 0 0 1 .668.279a.515.515 0 0 1-.277.671l-3.406 1.414l-.452 2.48H9.269l.257-1.426l-1.382.573a.51.51 0 0 1-.668-.279a.515.515 0 0 1 .277-.67l1.99-.826l.15-.839l-2.687 1.115a.51.51 0 0 1-.667-.279a.515.515 0 0 1 .277-.671l3.293-1.367L12.507 7.01h6.773c4.147-.143 6.22 1.242 6.22 4.155c0 3.695-2.702 6.553-7.285 6.553H12.94l-.283 1.556l1.487-.617a.51.51 0 0 1 .667.279a.515.515 0 0 1-.277.67l-2.097.871l-.153.84zm-.57-11.047l-1.172 6.425h5.218c3.212 0 4.43-2.185 4.43-3.77c0-1.584-.766-2.655-3.322-2.655z"/></g>`),
		g.Group(children),
	)
}