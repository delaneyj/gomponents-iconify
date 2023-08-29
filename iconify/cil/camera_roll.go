package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRoll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M328 104a32.036 32.036 0 0 0-32-32h-16V48a32.036 32.036 0 0 0-32-32H96a32.036 32.036 0 0 0-32 32v24H48a32.036 32.036 0 0 0-32 32v360a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32v-24h168V136H328Zm136 64v240H296v56H48V104h48V48h152v56h48v64Z"/><path fill="currentColor" d="M392 200h40v40h-40zm-72 0h40v40h-40zm-72 0h40v40h-40zm144 136h40v40h-40zm-72 0h40v40h-40zm-72 0h40v40h-40zm-72-136h40v40h-40zm0 136h40v40h-40z"/>`),
		g.Group(children),
	)
}