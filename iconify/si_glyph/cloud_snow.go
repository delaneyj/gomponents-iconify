package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 3)"><path d="M13.803 3.424c-.25 0-.486.067-.711.168c-.693-1.043-1.827-1.725-3.11-1.725c-.207 0-.407.021-.604.056C8.698.785 7.513.03 6.162.03c-1.907 0-3.483 1.501-3.771 3.462c-.018 0-.034-.005-.051-.005c-1.277 0-2.312 1.112-2.312 2.484h15.959c-.138-1.434-1.059-2.547-2.184-2.547z"/><circle cx="1.973" cy="7.973" r=".973"/><circle cx="5.962" cy="9.962" r=".962"/><circle cx="9.973" cy="7.973" r=".973"/><circle cx="13.962" cy="9.962" r=".962"/></g>`),
		g.Group(children),
	)
}