package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.545 2.1a.75.75 0 0 1 .274 1.025l-3.472 6.007a3 3 0 1 1-1.208-.908l1.759-3.042a6.5 6.5 0 0 0-2.148-.639V5a.75.75 0 1 1-1.5 0v-.457a6.5 6.5 0 0 0-1.829.49l.229.396a.75.75 0 1 1-1.3.75l-.228-.396a6.5 6.5 0 0 0-1.339 1.34l.396.227a.75.75 0 0 1-.75 1.3l-.396-.229a6.5 6.5 0 0 0-.498 1.905a.75.75 0 0 1-1.492-.155A8 8 0 0 1 11.65 3.88l.87-1.506a.75.75 0 0 1 1.025-.274Zm-.107 4.047a.75.75 0 0 1 1.047.169a8 8 0 0 1 1.51 4.963a.75.75 0 1 1-1.499-.052a6.5 6.5 0 0 0-1.227-4.033a.75.75 0 0 1 .17-1.047ZM9.5 11a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}