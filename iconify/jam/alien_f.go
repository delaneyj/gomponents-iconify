package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlienF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20.858c-4.5 0-9-6.03-9-11a9 9 0 1 1 18 0c0 4.97-4.5 11-9 11zm-1.974-6.181c.529-.192 1.42-1.946.853-3.503c-.566-1.557-2.399-2.32-2.905-2.135c-.506.184-1.42 1.946-.853 3.503c.566 1.557 2.376 2.327 2.905 2.135zm3.948 0c.529.192 2.339-.578 2.905-2.135c.567-1.557-.347-3.319-.853-3.503c-.506-.185-2.339.578-2.905 2.135c-.567 1.557.324 3.31.853 3.503z"/>`),
		g.Group(children),
	)
}