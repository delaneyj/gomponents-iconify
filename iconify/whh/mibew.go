package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mibew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 320V160L896 0l128 128v480zm64 512H512V512l160-128l288 256zM384 416V160l128-96l128 128v128L480 448zm64 384l-128 96L0 576l320-96l128 32v288zM0 64l320 96v256L0 512V64z"/>`),
		g.Group(children),
	)
}