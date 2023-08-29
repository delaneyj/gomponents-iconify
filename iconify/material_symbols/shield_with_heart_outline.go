package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldWithHeartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16q2.55-2.3 3.275-3.238T16 10.9q0-.9-.65-1.55T13.8 8.7q-.525 0-1.013.212T12 9.5q-.3-.375-.775-.588T10.2 8.7q-.9 0-1.55.65T8 10.9q0 .475.125.875t.55.938q.425.537 1.212 1.312T12 16Zm0 6q-3.475-.875-5.738-3.988T4 11.1V5l8-3l8 3v6.1q0 3.8-2.263 6.913T12 22Zm0-2.1q2.6-.825 4.3-3.3t1.7-5.5V6.375l-6-2.25l-6 2.25V11.1q0 3.025 1.7 5.5t4.3 3.3Zm0-7.9Z"/>`),
		g.Group(children),
	)
}