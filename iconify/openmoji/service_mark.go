package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" clip-rule="evenodd"><path stroke-width="6.581" d="M57.76 47.52V24.49l-9.871 19.74l-9.871-19.74v23.03"/><path stroke-width="6.574" d="M28.54 29.08c-.674-2.623-3.586-4.592-7.077-4.592c-3.991 0-7.225 2.574-7.225 5.753c0 3.175 3.235 5.753 7.225 5.753l-.148.02c3.991 0 7.225 2.577 7.225 5.753c0 3.179-3.235 5.753-7.225 5.753c-3.49 0-6.403-1.972-7.077-4.596"/></g>`),
		g.Group(children),
	)
}