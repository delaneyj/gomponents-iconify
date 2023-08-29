package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1223 768h660L603 2048H313l384-768H248L888 0h719l-384 768zM549 1920L1573 896h-557l384-768H967L455 1152h449l-384 768h29zm917-512h200l128 640h-130l-26-128h-285l-64 128h-143l320-640zm-49 384h196l-56-280l-140 280z"/>`),
		g.Group(children),
	)
}