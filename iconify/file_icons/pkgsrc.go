package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pkgsrc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 143.184L21.893 400.85L241.692 512l-2.552-272.82L0 143.184zm294.717 92.605l191.986-114.517l-20.21 252.615L291.756 510.02l2.96-274.232zm-32.002-48.853l195.35-109.467L227.349 0L21.89 94.31l240.824 92.626z"/>`),
		g.Group(children),
	)
}