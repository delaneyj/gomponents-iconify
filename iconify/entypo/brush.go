package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.763 13.563c-1.515 1.488-.235 3.016-2.247 5.279c-.908 1.023 3.738.711 6.039-1.551c.977-.961.701-2.359-.346-3.389c-1.047-1.028-2.47-1.3-3.446-.339zM19.539.659C18.763-.105 10.16 6.788 7.6 9.305c-1.271 1.25-1.695 1.92-2.084 2.42c-.17.219.055.285.154.336c.504.258.856.496 1.311.943c.456.447.699.793.959 1.289c.053.098.121.318.342.152c.51-.383 1.191-.801 2.462-2.049C13.305 9.88 20.317 1.422 19.539.659z"/>`),
		g.Group(children),
	)
}