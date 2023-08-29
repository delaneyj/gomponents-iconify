package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Callalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m884 884l-91-91q49-53 76-126t27-155t-27-155t-76-126l91-91q66 72 103 168t37 204t-37 204t-103 168zM748 748l-90-90q46-58 46-146t-46-146l90-90q40 45 62 106t22 130t-22 130t-62 106zm-300 276h-64q-80 0-149.5-35.5t-122-100t-82.5-162T0 512t30-214.5t82.5-162t122-100T384 0h64q27 0 45.5 18.5T512 64v192q0 26-18.5 45T448 320h-64q-24 0-42-15.5T321 265q-129 40-129 247t129 247q3-24 21-39.5t42-15.5h64q26 0 45 18.5t19 45.5v192q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}