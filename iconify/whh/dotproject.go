package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dotproject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M864 320q-66 0-113-47t-47-113t47-113T864 0t113 47t47 113t-47 113t-113 47zM384 0h256v1024H384V0zM160 1024q-66 0-113-47T0 864t47-113t113-47t113 47t47 113t-47 113t-113 47z"/>`),
		g.Group(children),
	)
}