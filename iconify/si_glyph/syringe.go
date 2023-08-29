package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syringe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="m.032 4.258l.647.648L4.892.693L4.245.046L.032 4.258z"/><path fill="currentColor" d="M1.844 3.207L3.45 4.813l1.4-1.4l-1.606-1.605l-1.4 1.4z"/><path d="M5.931 5.487h1.104v6.062H5.931zm1.917-2.366h1.105v6.062H7.848z"/><path fill="currentColor" d="m13.548 15.369l1.749-.928l-2.629-2.731l-1.391 1.39l2.271 2.269zM6.79 1.974L1.633 7.13c-.131.131-.154.452-.02.585c.131.133.494.283.626.152l.513-.514l5.672 5.671l-.014.014l.671.673l.68-.68l.741.742l2.799-2.8l-.741-.741l.722-.721l-.672-.672l-.016.015l-5.671-5.672l.47-.47c.131-.132-.029-.504-.162-.637c-.131-.132-.309-.229-.441-.101zM5.493 4.581l.781-.78l4.287 4.286l-.781.781l-4.287-4.287zM3.372 6.703l.781-.781l4.286 4.287l-.78.78l-4.287-4.286zm5.687 5.687l-.75-.751l2.888-2.888l.751.751l-2.889 2.888z"/></g>`),
		g.Group(children),
	)
}