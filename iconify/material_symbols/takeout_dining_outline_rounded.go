package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeoutDiningOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.8 18h8.45l.475-7H7.3l.5 7Zm-.65-9h9.725l.075-1.25L14.15 5h-4.3l-2.8 2.75l.1 1.25Zm-1.9 1.7L2.7 8.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l.9.9l-.05-.6l3.475-3.475q.275-.275.638-.425T9.825 3h4.35q.4 0 .763.15t.637.425L19.05 7.05l-.05.6l.9-.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-2.55 2.55H5.25ZM7.8 20q-.8 0-1.375-.525T5.8 18.15l-.55-7.45h13.5l-.55 7.45q-.05.8-.625 1.325T16.2 20H7.8ZM12 9Zm.025 2Z"/>`),
		g.Group(children),
	)
}