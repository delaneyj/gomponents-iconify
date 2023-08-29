package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Failed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m557.248 608l135.744-135.744l-45.248-45.248l-135.68 135.744l-135.808-135.68l-45.248 45.184L466.752 608l-135.68 135.68l45.184 45.312L512 653.248l135.744 135.744l45.248-45.248L557.312 608zM704 192h160v736H160V192h160v64h384v-64zm-320 0V96h256v96H384z"/>`),
		g.Group(children),
	)
}