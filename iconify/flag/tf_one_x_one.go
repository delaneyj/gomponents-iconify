package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TfOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><path id="flagTf1x10" fill="#fff" d="m0-21l12.3 38L-20-6.5h40L-12.3 17z"/></defs><path fill="#002395" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 0h312.3v210H0z"/><path fill="#002395" d="M0 0h102.4v204.8H0z"/><path fill="#ed2939" d="M204.8 0h102.4v204.8H204.8z"/><path fill="#fff" d="m282.4 234.2l16.5 26.3h46.9V352l-35.3-55l-47.3 75.5h23l24.4-43.5l49.9 89.6l49.9-89.6l24.3 43.5h23L410.5 297l-35.2 55v-50.6h21.1l15.7-25h-36.8v-16h46.9l16.5-26.2H282.4zm55 112h-51.2v18h51.2zm97.3 0h-51.2v18h51.2z"/><use width="100%" height="100%" x="416" y="362" href="#flagTf1x10" transform="translate(-172) scale(1.28)"/><use width="100%" height="100%" x="371" y="328" href="#flagTf1x10" transform="translate(-172) scale(1.28)"/><use width="100%" height="100%" x="461" y="328" href="#flagTf1x10" transform="translate(-172) scale(1.28)"/><use width="100%" height="100%" x="333" y="227" href="#flagTf1x10" transform="translate(-172) scale(1.28)"/><use width="100%" height="100%" x="499" y="227" href="#flagTf1x10" transform="translate(-172) scale(1.28)"/>`),
		g.Group(children),
	)
}