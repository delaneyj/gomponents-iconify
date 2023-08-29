package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buttonb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm166-546q26-45 26-94q0-80-56-136t-136-56H320q-26 0-45 19t-19 45v512q0 26 19 45t45 19h256q80 0 136-56t56-136q0-50-24-93t-66-69zM576 704H384V576h192q27 0 45.5 19t18.5 45.5t-19 45t-45 18.5zm-64-256H384V320h128q26 0 45 19t19 45.5t-18.5 45T512 448z"/>`),
		g.Group(children),
	)
}