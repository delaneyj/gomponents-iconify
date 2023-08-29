package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Okx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.167 17.833H17.833v12.334h12.334V17.833ZM42.5 30.167H30.167V42.5H42.5V30.167Zm0-24.667H30.167v12.333H42.5V5.5ZM17.833 30.167H5.5V42.5h12.333V30.167Zm0-24.667H5.5v12.333h12.333V5.5Z"/>`),
		g.Group(children),
	)
}