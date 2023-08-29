package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#2980BA" d="M50 0c27.613 0 50 22.386 50 50s-22.387 50-50 50C22.386 100 0 77.614 0 50S22.386 0 50 0z"/><defs><circle id="flatUiShirt0" cx="50" cy="50" r="50"/></defs><clipPath id="flatUiShirt1"><use href="#flatUiShirt0"/></clipPath><g clip-path="url(#flatUiShirt1)"><path fill="#2D3E50" d="M-15 17h133v87H-15V17z"/><path fill="#EDF0F1" d="M51 17v95h32V17H51zm-34 95h32V17H17v95z"/><path fill="#E64B3C" d="M73 26H27l18.095 21.905L38 55v48h24V55l-7.098-7.098z"/><path fill="#D5D9D8" d="m26 17l.049 36L50 29l.082-12H26z"/><path fill="#D5D9D8" d="m74 17l-.049 36l-23.858-24l-.081-12H74z"/><path fill="#C03A2B" d="M73 26H27l11.771 14.251l11.274-11.295l11.203 11.266z"/><path fill="#fff" d="M69 12v5H31v-5l-5 5v33l24-24l24 24V17z"/><path fill="#C03A2B" d="m56 49l-1.098-1.098l.744-.902H44.347l.748.905L44 49z"/><path fill="#DE9273" d="M31-1v19.674h.045c.241 1.587 2.065 3.15 5.517 4.362c7.42 2.606 19.452 2.606 26.872 0c3.45-1.212 5.274-2.775 5.518-4.362H69V-1H31z"/><path fill="#BE7C63" d="M69-1s.014 4.435-.018 4.467A17.943 17.943 0 0 1 56 9H44a17.948 17.948 0 0 1-13.022-5.573C31.002 3.452 31-1 31-1h38z"/><path fill="#E39D82" d="M50-1v25.991c4.861 0 9.725-.652 13.434-1.955c3.45-1.212 5.274-2.775 5.518-4.362H69V-1H50z"/><path fill="#C0866E" d="M56 9c5.102 0 9.708-2.123 12.982-5.533c.01-.009.015-.395.018-.93V-1H50V9h6z"/></g>`),
		g.Group(children),
	)
}