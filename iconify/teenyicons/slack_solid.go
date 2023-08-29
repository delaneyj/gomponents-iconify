package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlackSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10.385 6.923H8.077v-5.77a1.154 1.154 0 0 1 2.308 0v5.77Zm2.307 0H11.54V5.77a1.154 1.154 0 1 1 1.153 1.154Zm1.154 1.154h-5.77v2.308h5.77a1.154 1.154 0 0 0 0-2.308Zm-5.769 4.615V11.54H9.23a1.154 1.154 0 1 1-1.154 1.153ZM1.154 4.615h5.77v2.308h-5.77a1.154 1.154 0 0 1 0-2.308Zm5.769-2.307v1.154H5.77a1.154 1.154 0 1 1 1.154-1.154ZM1.154 9.23c0-.636.516-1.153 1.154-1.153h1.154V9.23a1.154 1.154 0 0 1-2.308 0Zm3.461 4.616v-5.77h2.308v5.77a1.154 1.154 0 0 1-2.308 0Z"/>`),
		g.Group(children),
	)
}