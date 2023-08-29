package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeKotlin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><linearGradient id="vscodeIconsFileTypeKotlin0" x1="311.336" x2="283.342" y1="1452.064" y2="1480.058" gradientTransform="translate(-281.4 -1450)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#e44857"/><stop offset=".47" stop-color="#9d4b9d"/><stop offset="1" stop-color="#6d5faa"/></linearGradient></defs><path fill="url(#vscodeIconsFileTypeKotlin0)" d="M30 30H2V2h28L16 16Z"/>`),
		g.Group(children),
	)
}