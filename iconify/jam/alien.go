package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20c-4.5 0-9-6.03-9-11a9 9 0 1 1 18 0c0 4.97-4.5 11-9 11zm0-2c3.14 0 7-4.869 7-9A7 7 0 0 0 2 9c0 4.131 3.86 9 7 9zm-1.974-4.18c-.529.192-2.339-.579-2.905-2.136c-.567-1.557.347-3.319.853-3.503c.506-.184 2.339.578 2.905 2.135c.567 1.557-.324 3.31-.853 3.503zm3.948 0c-.529-.193-1.42-1.947-.853-3.504c.566-1.557 2.399-2.32 2.905-2.135c.506.184 1.42 1.946.853 3.503c-.566 1.557-2.376 2.328-2.905 2.135z"/>`),
		g.Group(children),
	)
}