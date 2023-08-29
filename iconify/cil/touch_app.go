package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M454.423 278.957L328 243.839v-8.185a116 116 0 1 0-104 0V312h-24.418l-18.494-22.6a90.414 90.414 0 0 0-126.43-13.367a20.862 20.862 0 0 0-8.026 33.47L215.084 496H472V302.08a24.067 24.067 0 0 0-17.577-23.123ZM192 132a84 84 0 1 1 136 65.9V132a52 52 0 0 0-104 0v65.9a83.866 83.866 0 0 1-32-65.9Zm248 332H229.3L79.141 297.75a58.438 58.438 0 0 1 77.181 11.91l28.1 34.34H256V132a20 20 0 0 1 40 0v136.161l144 40Z"/>`),
		g.Group(children),
	)
}