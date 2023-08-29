package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freeway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 448h-32v160q0 13-9.5 22.5T928 640t-22.5-9.5T896 608V448H128v160q0 13-9.5 22.5T96 640t-22.5-9.5T64 608V448H32q-13 0-22.5-9.5T0 416v-64q0-13 9.5-22.5T32 320h960q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 448zM576 32q0-13 9.5-22.5T608 0h64q13 0 22.5 9.5T704 32l45 224H576V32zm-256 0q0-13 9.5-22.5T352 0h64q13 0 22.5 9.5T448 32v224H275zm128 480v480q0 13-9.5 22.5T416 1024H160q-13 0-22.5-9.5T128 992l96-480h224zm448 480q0 13-9.5 22.5T864 1024H608q-13 0-22.5-9.5T576 992V512h224z"/>`),
		g.Group(children),
	)
}