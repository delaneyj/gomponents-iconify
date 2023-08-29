package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockUnlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.927 0c-2.148 0-3.896 1.802-3.896 4.017v1.662a5.403 5.403 0 0 0-2.552-.651a5.457 5.457 0 0 0-5.458 5.458a5.458 5.458 0 1 0 10.916 0a5.38 5.38 0 0 0-1-3.106V4.018c0-1.112.943-2.017 1.989-2.017h.177c1.046 0 1.896.904 1.896 2.017v2.883h2V4.018C17 1.802 15.252 0 13.104 0h-.177zM7.416 11.467l2.578-.753a3.528 3.528 0 0 1-3.515 3.312a3.539 3.539 0 1 1 0-7.076A3.532 3.532 0 0 1 10 10.299l-2.573-.781a1.352 1.352 0 0 0-.947-.388a1.358 1.358 0 0 0-.001 2.716c.363 0 .693-.147.937-.379z"/>`),
		g.Group(children),
	)
}