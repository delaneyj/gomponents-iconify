package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM242.285 329.297h715.356v541.406H242.285V329.297zm87.744 87.744v365.918H869.97V417.041H330.029zm189.698 63.721l210.719 121.51l-210.719 122.24v-243.75z"/>`),
		g.Group(children),
	)
}