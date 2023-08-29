package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatboxEllipses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M408 48H104a72.08 72.08 0 0 0-72 72v192a72.08 72.08 0 0 0 72 72h24v64a16 16 0 0 0 26.25 12.29L245.74 384H408a72.08 72.08 0 0 0 72-72V120a72.08 72.08 0 0 0-72-72ZM160 248a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}