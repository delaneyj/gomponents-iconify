package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmassTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m6.781 3.981l1.221-.936l1.217.936l-.465-1.516l1.219-.938H8.467L8.002.011l-.467 1.516H6.028l1.22.938l-.467 1.516z"/><path d="m10.984 11.979l1.934.016l-2.902-4.031H12L7.985 3.98L4.031 7.996h2.031l-3.047 3.953h2.016l-3.016 4.016l11.977-.017l-3.008-3.969zm-4.968 3.032H4.969v-1.062h1.047v1.062zm1-4H5.969V9.949h1.047v1.062zm1.015-4H6.969V5.949h1.062v1.062zm.938 2.937h1.047v1.062H8.969V9.948zm1 5.063v-1.062h1.047v1.062H9.969z"/><path d="M12 7.995h.969v.984H12zm1 4h.969v.984H13zm-11 0h.969v.984H2zm1-4h.969v.984H3z"/></g>`),
		g.Group(children),
	)
}