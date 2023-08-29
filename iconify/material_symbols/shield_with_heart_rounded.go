package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldWithHeartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10.9q0 1.4 1.375 2.675t2.3 2.125q.15.125.338.125t.337-.125q.9-.825 2.275-2.137T16 10.9q0-.9-.65-1.55T13.8 8.7q-.525 0-1.013.213T12 9.5q-.3-.375-.775-.588T10.2 8.7q-.9 0-1.55.65T8 10.9Zm4 11q-.175 0-.325-.025t-.3-.075Q8 20.675 6 17.637T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 3.5-2 6.538T12.625 21.8q-.15.05-.3.075T12 21.9Z"/>`),
		g.Group(children),
	)
}