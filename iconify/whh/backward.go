package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M962 1014L527 551q-15-16-15-38.5t15-37.5L962 11q25-26 62 14v977q-37 39-62 12zm-512 0L15 551Q0 535 0 512.5T15 475L450 11q25-26 62 14v977q-37 39-62 12z"/>`),
		g.Group(children),
	)
}