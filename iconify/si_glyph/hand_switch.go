package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.484 3.062c-.82 0-1.484.637-1.484 1.422c0 .062.011.121.02.181L9.317 6.853c-.739-1.108-1.866-1.848-3.348-1.848l-.016.001l-.007-2.604c0-.715-.775-1.294-1.525-1.316c-.751.022-1.353.602-1.353 1.316v13.13c0 .715.602 1.294 1.353 1.316c.75-.022 1.525-.602 1.525-1.316v-2.604l.022.002c2.3 0 3.993-1.775 3.993-3.963c0-.309-.04-.607-.11-.895l3.955-2.332c.204.102.432.164.678.164c.82 0 1.484-.637 1.484-1.422c0-.785-.663-1.42-1.484-1.42z"/>`),
		g.Group(children),
	)
}