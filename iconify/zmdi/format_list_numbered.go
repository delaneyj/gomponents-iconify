package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 299v-22h64v86H0v-22h43v-10H21v-22h22v-10H0zm21-192V43H0V21h43v86H21zM0 171v-22h64v20l-38 44h38v22H0v-20l38-44H0zM107 43h298v42H107V43zm0 298v-42h298v42H107zm0-128v-42h298v42H107z"/>`),
		g.Group(children),
	)
}