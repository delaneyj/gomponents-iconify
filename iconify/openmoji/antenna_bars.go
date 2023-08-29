package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiAntennaBars0" d="M12 51h5v9h-5zm10.75-10h5v19h-5zM33.5 31h5v28.999h-5zm10.75-10h5v38.999h-5zM55 12h5v48h-5z"/></defs><use href="#openmojiAntennaBars0" fill="#d0cfce" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><use href="#openmojiAntennaBars0" fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}