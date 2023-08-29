package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capacitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#53b9ff" d="M19.93 27.059L.156 46.859l30.496 30.59L0 108.191l19.715 19.813L50.43 97.25l30.547 30.531l19.777-19.8Zm0 0"/><path fill="#119eff" d="M70.258 77.45L50.43 97.25l30.547 30.531l19.777-19.8Zm0 0"/><path fill-opacity=".2" d="M70.258 77.45L50.43 97.25l7.633 7.59Zm0 0"/><path fill="#53b9ff" d="M97.285 50.492L128 19.738L108.215 0L77.512 30.691L46.957.156L27.184 19.957l80.82 80.922l19.777-19.8Zm0 0"/><path fill="#119eff" d="m57.68 50.492l19.828-19.8L46.957.155L27.184 19.957Zm0 0"/><path fill-opacity=".2" d="m57.68 50.492l19.828-19.8l-7.633-7.594Zm0 0"/>`),
		g.Group(children),
	)
}