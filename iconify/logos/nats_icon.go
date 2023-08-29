package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#34A574" d="M128 0h128v103.768H128z"/><path fill="#27AAE1" d="M0 0h128v103.768H0z"/><path fill="#8DC63F" d="M256 103.863v103.769h-84.193v57.395l-62.622-57.205l18.815-.76V103.863z"/><path fill="#375C93" d="M128 103.863v120.678l-18.815-16.719H0V103.863z"/><path fill="#FFF" d="M181.024 134.177V48.273h30.599v111.086H165.25l-93.6-87.424v87.519H40.956V48.273h47.988l92.08 85.904z"/>`),
		g.Group(children),
	)
}