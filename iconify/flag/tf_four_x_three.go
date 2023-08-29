package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TfFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><path id="flagTf4x30" fill="#fff" d="m0-21l12.3 38L-20-6.5h40L-12.3 17z"/></defs><path fill="#002395" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0h292.8v196.8H0z"/><path fill="#002395" d="M0 0h96v192H0z"/><path fill="#ed2939" d="M192 0h96v192h-96z"/><path fill="#fff" d="m426 219.6l15.4 24.6h44V330l-33-51.6l-44.4 70.8h21.6l22.8-40.8l46.8 84l46.8-84l22.8 40.8h21.6L546 278.4L513 330v-47.4h19.8l14.7-23.4H513v-15h44l15.4-24.6H426zm51.6 105h-48v16.8h48zm91.2 0h-48v16.8h48z"/><use width="100%" height="100%" x="416" y="362" href="#flagTf4x30" transform="scale(1.2)"/><use width="100%" height="100%" x="371" y="328" href="#flagTf4x30" transform="scale(1.2)"/><use width="100%" height="100%" x="461" y="328" href="#flagTf4x30" transform="scale(1.2)"/><use width="100%" height="100%" x="333" y="227" href="#flagTf4x30" transform="scale(1.2)"/><use width="100%" height="100%" x="499" y="227" href="#flagTf4x30" transform="scale(1.2)"/>`),
		g.Group(children),
	)
}