package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeVscodeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#007acc" d="M29.821 4.321L24.023 2l-12.53 12.212l-7.66-5.827l-1.654.837V22.8l1.644.827l7.65-5.827L24.023 30l5.8-2.321V4.321ZM4.65 19.192v-6.374l3.55 3.167l-3.55 3.207ZM16 15.985l7.082-5.3v10.639l-7.092-5.339H16Z"/>`),
		g.Group(children),
	)
}