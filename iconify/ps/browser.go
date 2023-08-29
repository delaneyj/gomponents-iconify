package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 3H43Q25 3 12.5 15.5T0 45v363q0 21 21 21h470q21 0 21-21V45q0-17-12.5-29.5T469 3zm0 384H43V152h426v235zm0-278H43V45h426v64zm-277 86H85v149h107V195zm-43 106h-21v-64h21v64zm86-106h192v42H235v-42zm0 85h149v43H235v-43z"/>`),
		g.Group(children),
	)
}