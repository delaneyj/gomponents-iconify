package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M200 48H72a24.028 24.028 0 0 0-24 24v368a24.028 24.028 0 0 0 24 24h128a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24Zm-8 384H80V80h112ZM440 48H312a24.028 24.028 0 0 0-24 24v368a24.028 24.028 0 0 0 24 24h128a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24Zm-8 384H320V80h112Z"/>`),
		g.Group(children),
	)
}