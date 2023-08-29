package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftWindowsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#0078D4" d="M0 0h121.329v121.329H0V0Zm134.671 0H256v121.329H134.671V0ZM0 134.671h121.329V256H0V134.671Zm134.671 0H256V256H134.671V134.671Z"/>`),
		g.Group(children),
	)
}