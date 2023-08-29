package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HammerSledgeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M10.033.555a2 2 0 0 1 1.225.43l6.863 5.416a2 2 0 0 1 .703 1.091l.19.774a2 2 0 0 1-.371 1.717l-2.492 3.158a2 2 0 0 1-1.584.76l-.796-.005a2 2 0 0 1-1.225-.43L5.683 8.05a2 2 0 0 1-.703-1.09l-.19-.774a2 2 0 0 1 .371-1.718l2.492-3.157A2 2 0 0 1 9.237.55l.796.005Zm6.85 7.416l-6.864-5.416l-.796-.005L6.73 5.707l.19.773l6.864 5.416l.796.006l2.492-3.158l-.19-.773Z"/><path d="m9.028 10.061l-5.262 6.668l.785.62l5.262-6.669l1.57 1.24l-5.262 6.667a2 2 0 0 1-2.809.331l-.785-.62a2 2 0 0 1-.33-2.808l5.26-6.668l1.57 1.239Zm4.491-5.691l.93-1.178l.784.62l-.93 1.177l1.571 1.239l.93-1.178a2 2 0 0 0-.332-2.809l-.785-.619a2 2 0 0 0-2.809.331l-.93 1.178l1.57 1.239Z"/></g><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}