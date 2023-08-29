package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PantoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPantoneFill0"><g id="evaPantoneFill1"><path id="evaPantoneFill2" fill="currentColor" d="M20 13.18h-2.7l-1.86 2L11.88 19l-1.41 1.52L10 21h10a1 1 0 0 0 1-1v-5.82a1 1 0 0 0-1-1ZM18.19 9.3l-4.14-3.86a.89.89 0 0 0-.71-.26a1 1 0 0 0-.7.31l-.82.89v10.71a5.23 5.23 0 0 1-.06.57l6.48-6.95a1 1 0 0 0-.05-1.41ZM10.82 4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v13.09a3.91 3.91 0 0 0 7.82 0Zm-2 13.09a1.91 1.91 0 0 1-3.82 0V15h3.82Zm0-4.09H5v-3h3.82Zm0-5H5V5h3.82Z"/></g></g>`),
		g.Group(children),
	)
}