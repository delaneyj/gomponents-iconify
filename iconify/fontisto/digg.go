package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.84 8.364V4.802h2.394v11.77H0V8.365zm0 6.282v-4.359H2.39v4.359zm13.92-6.282H24v11.44h-6.234v-1.92h3.84v-1.313h-3.84zm3.84 6.282v-4.359h-1.44v4.359zm-11.062 1.92V8.364h6.234v11.44h-6.225v-1.92h3.84v-1.313zm2.39-6.282v4.359h1.453v-4.359zM9.6 4.8v2.39H7.172V4.8zm0 3.562v8.209H7.172V8.364z"/>`),
		g.Group(children),
	)
}