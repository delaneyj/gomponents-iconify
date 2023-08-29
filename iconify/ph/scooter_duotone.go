package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScooterDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M72 172a28 28 0 1 1-28-28a28 28 0 0 1 28 28Zm140-28a28 28 0 1 0 28 28a28 28 0 0 0-28-28Z" opacity=".2"/><path d="M212 136c-1.18 0-2.35.06-3.51.17l-10.75-32.25l-22.15-66.45A8 8 0 0 0 168 32h-32a8 8 0 0 0 0 16h26.23l19 56.87L132.09 168H79.77a36 36 0 1 0-1.83 16H136a8 8 0 0 0 6.31-3.09l45-57.8l6 18.13A36 36 0 1 0 212 136ZM44 192a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm168 0a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/></g>`),
		g.Group(children),
	)
}