package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 1025V0h1024v1025H0zM960 64H64v769h896V64zm-704 64q53 0 90.5 37.5T384 256t-37.5 90.5T256 384t-90.5-37t-37.5-90v-1q0-53 37.5-90.5T256 128zm111 402l238 239H128v-93l146-146q18-18 46.5-18t46.5 18zm323-192q18-19 46.5-19t46.5 19l113 113v318H692q0-1-1.5-3.5T688 761L477 551z"/>`),
		g.Group(children),
	)
}