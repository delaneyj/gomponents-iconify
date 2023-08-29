package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirroring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 39.892a2.323 2.323 0 0 1-2.329 2.329h-4.075a2.323 2.323 0 0 1-2.328-2.329v-4.075c0-10.39-9.062-21.306-21.864-21.306H7.829A2.323 2.323 0 0 1 5.5 12.183V8.108a2.323 2.323 0 0 1 2.33-2.329h19.355c10.629 0 15.274 4.771 15.316 13.796v20.317Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.558 39.892a2.323 2.323 0 0 1-2.328 2.329h-4.075a2.323 2.323 0 0 1-2.329-2.329v-2.568c0-5.942-4.533-9.673-10.183-9.673H7.83a2.323 2.323 0 0 1-2.33-2.329v-4.075a2.323 2.323 0 0 1 2.33-2.327h1.814c11.23 0 18.916 7.392 18.916 18.154v2.82Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.829 33.489h1.395c3.433 0 5.008 1.952 5.008 5.05v1.353a2.323 2.323 0 0 1-2.328 2.329H7.829A2.323 2.323 0 0 1 5.5 39.892v-4.075a2.323 2.323 0 0 1 2.329-2.328Z"/>`),
		g.Group(children),
	)
}