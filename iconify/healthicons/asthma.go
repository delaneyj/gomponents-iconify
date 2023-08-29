package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asthma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 6c-.364 0-.725.015-1.082.044A1 1 0 0 0 5 7.041V43a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-4.483l4.097 1.755c.524.225 1.16.461 1.825.549c.677.09 1.446.031 2.16-.405A3.998 3.998 0 0 0 20 37v-3a1 1 0 0 0-1-1h-3a2 2 0 0 1 2-2h1a1 1 0 0 0 1-1v-2h1.131c1.598 0 2.55-1.78 1.665-3.11L20 20.698V19c0-7.18-5.82-13-13-13Zm10 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm18.497-5.76a1 1 0 0 1 1.223-.706l4.468 1.192a1 1 0 0 1 .708 1.226l-.777 2.888l.294.133a1 1 0 0 1 .551 1.176l-2.94 10.678A5 5 0 0 1 34.205 36H28a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h3.247l2.385-8.78a1 1 0 0 1 1.372-.65l.493-1.83Zm4.208 2.16l-.463 1.72l-2.422-1.102l.349-1.294l2.536.676Z" clip-rule="evenodd"/><path d="M26 28v7h-2v-7h2Z"/></g>`),
		g.Group(children),
	)
}