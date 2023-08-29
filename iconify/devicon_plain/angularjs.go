package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angularjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M52.864 64h23.28L63.769 38.123zM63.81 1.026L4.553 21.88l9.363 77.637l49.957 27.457l50.214-27.828l9.36-77.635L63.81 1.026zM48.044 75l-7.265 18.176l-13.581.056l36.608-81.079l-.07-.153h-.064l.001-.133l.063.133h.141l.123-.274V12h-.124l-.069.153l38.189 81.417l-13.074-.287l-8.042-18.58l-17.173.082"/>`),
		g.Group(children),
	)
}