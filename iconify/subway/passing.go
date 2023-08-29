package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M186.2 232.7h139.6v-46.5H186.2v46.5zm0 93.1h139.6v-46.5H186.2v46.5zm0 93.1h139.6v-46.5H186.2v46.5zm0 93.1h139.6v-46.5H186.2V512zm186.2-93.1H512v-46.5H372.4v46.5zm0 93.1H512v-46.5H372.4V512zm0-279.3H512v-46.5H372.4v46.5zm0-139.6v46.5H512V93.1H372.4zm0 232.7H512v-46.5H372.4v46.5zM0 232.7h139.6v-46.5H0v46.5zm0 93.1h139.6v-46.5H0v46.5zm0 93.1h139.6v-46.5H0v46.5zM0 46.5h139.6V0H0v46.5zm0 93.1h139.6V93.1H0v46.5zM0 512h139.6v-46.5H0V512z"/>`),
		g.Group(children),
	)
}