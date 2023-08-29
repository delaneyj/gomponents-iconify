package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1243 1920h421v128H677l-248-248q-27-27-41-62t-15-74q0-38 14-73t42-63l435-435l-65-197H353l-85 256H128L512 0h128l329 988l375-374l602 602l-703 704zM756 768L576 228L396 768h360zm588 26l-550 550l422 422l550-550l-422-422zm-283 1126l65-64l-422-422l-184 185q-19 19-19 45t19 45l211 211h330z"/>`),
		g.Group(children),
	)
}