package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerSledgeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHammerSledgeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.28 4.055a1.5 1.5 0 0 1 .918.323l6.863 5.415a1.5 1.5 0 0 1 .528.819l.19.773a1.5 1.5 0 0 1-.279 1.289l-2.49 3.156a1.5 1.5 0 0 1-1.188.57l-.797-.005a1.5 1.5 0 0 1-.918-.322l-6.863-5.416a1.5 1.5 0 0 1-.528-.819l-.19-.773a1.5 1.5 0 0 1 .279-1.288l2.491-3.157a1.5 1.5 0 0 1 1.188-.571l.797.005Zm.299 1.108a.5.5 0 0 0-.306-.108l-.797-.005a.5.5 0 0 0-.396.19L9.59 8.397a.5.5 0 0 0-.093.43l.19.773a.5.5 0 0 0 .176.273l6.863 5.416a.5.5 0 0 0 .306.107l.797.006a.5.5 0 0 0 .396-.19l2.491-3.158a.5.5 0 0 0 .093-.43l-.19-.773a.5.5 0 0 0-.176-.273L13.58 5.163Zm-1.694 7.588L6.624 19.42a.5.5 0 0 0 .082.702l.785.62a.5.5 0 0 0 .703-.083l5.261-6.668l.785.62l-5.261 6.668a1.5 1.5 0 0 1-2.107.248l-.785-.62a1.5 1.5 0 0 1-.248-2.106l5.261-6.668l.785.62Z"/><path d="m16.376 7.06l.93-1.178a.5.5 0 0 1 .702-.082l.785.62a.5.5 0 0 1 .083.701l-.93 1.178l.785.62l.93-1.178a1.5 1.5 0 0 0-.249-2.107l-.785-.62a1.5 1.5 0 0 0-2.106.249l-.93 1.177l.785.62Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHammerSledgeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}