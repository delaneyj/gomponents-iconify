package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationAngledownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21q-.425 0-.713-.3T10 19.975q0-.425.288-.713t.712-.287h.6L2.975 10.35q-.3-.3-.3-.712t.3-.713q.3-.3.7-.3t.7.3l8.65 8.65v-.6q0-.425.275-.713t.7-.287q.425 0 .713.288t.287.712v3q0 .425-.3.725t-.725.3H11Zm3.175-6.475q-.175-.175-.225-.438t.075-.487l1.05-2.15l-3.15-3.125l-2.15 1q-.225.125-.488.075t-.462-.25q-.375-.375-.312-.8t.562-.65l8.6-3.925q.375-.175.85-.088t.775.388q.275.275.363.75t-.088.85l-3.925 8.6q-.225.5-.65.575t-.825-.325Zm1.6-4.625l2.1-4.35l-.075-.075l-4.35 2.1L15.775 9.9Z"/>`),
		g.Group(children),
	)
}