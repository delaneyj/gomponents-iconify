package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Limbo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.804 44.5c-.158-2.108-.764-4.953-2.872-6.35v-1.843c2.925-1.45 7.482-6.481 7.482-9.906c0-5.4-1.976-10.222-2.924-12.25c3.767-1.803 3.909-4.604 2.292-3.53c3.27-1.767 2.305-3.456 1.185-2.793c.353-1.103-.976-2.123-2.239-1.66c.755-2.081-3.255-1.9-4.479-1.238c.159-1.818-.491-2.3-7.113.948c.462-1.77-.651-1.822-5.269 1.423c-2.502.553-2.977 1.95-2.977 1.95c-3.887-.47-8.162 5.705-7.745 8.483c-1.554.983-1.475 2.476-.237 1.817c-.685 6.903 3.714 10.538 8.088 13.2c4.373 2.66 5.717 3.53 6.586 4.662c-1.739 1.818-3.32 3.9-3.952 7.087"/><circle cx="27.964" cy="23.424" r=".75" fill="currentColor"/><circle cx="34.94" cy="21.062" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}