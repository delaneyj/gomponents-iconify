package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosKeypad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M394.6 341.2c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M256 341.2c-29.5 0-53.4 23.9-53.4 53.4S226.5 448 256 448s53.4-23.9 53.4-53.4-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M117.4 341.2c-29.5 0-53.4 23.9-53.4 53.4S87.9 448 117.4 448s53.4-23.9 53.4-53.4-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M394.6 202.6c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4S448 285.5 448 256s-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M256 202.6c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M117.4 202.6C87.9 202.6 64 226.5 64 256s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4-23.9-53.4-53.4-53.4z" fill="currentColor"/><path d="M394.6 64c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4S424.1 64 394.6 64z" fill="currentColor"/><path d="M256 64c-29.5 0-53.4 23.9-53.4 53.4s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4S285.5 64 256 64z" fill="currentColor"/><path d="M117.4 64C87.9 64 64 87.9 64 117.4s23.9 53.4 53.4 53.4 53.4-23.9 53.4-53.4S146.9 64 117.4 64z" fill="currentColor"/>`),
		g.Group(children),
	)
}