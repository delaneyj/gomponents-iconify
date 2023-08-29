package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M124 136H36a20.023 20.023 0 0 0-20 20v320a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V156a20.023 20.023 0 0 0-20-20Zm-12 328H48V168h64Zm188-224h-88a20.023 20.023 0 0 0-20 20v216a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V260a20.023 20.023 0 0 0-20-20Zm-12 224h-64V272h64ZM476 16h-88a20.023 20.023 0 0 0-20 20v440a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V36a20.023 20.023 0 0 0-20-20Zm-12 448h-64V48h64Z"/>`),
		g.Group(children),
	)
}